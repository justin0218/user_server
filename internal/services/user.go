package services

import (
	"user_server/api/auth_server"
	"user_server/api/proto"
	"user_server/internal/models/user"
	"user_server/pkg/email"
	"context"
	"fmt"
	"math/rand"
	"time"
)

//
type UserService struct {
	baseService
}

//
func (s *UserService) SendEmailCode(emailStr, from string) (err error) {
	redisClient := s.Redis.Get()
	rkey := s.Redis.GetSendEmailKey(emailStr, from)
	code, _ := redisClient.Get(rkey).Result()
	if code != "" {
		err = fmt.Errorf("操作频繁，请稍后再试")
		return
	}
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	code = fmt.Sprintf("%06v", rnd.Int31n(1000000))
	con := fmt.Sprintf("欢迎注册己课堂，您的验证码：%s", code)
	err = email.Send(emailStr, con, "己课堂注册验证码")
	if err != nil {
		return
	}
	err = redisClient.Set(rkey, code, time.Second*60).Err()
	return
}

//
func (s *UserService) Register(emailStr, vcode string) (userFull *proto.RegisterRes, err error) {
	userFull = new(proto.RegisterRes)
	ctx := context.Background()
	redisClient := s.Redis.Get()
	rkey := s.Redis.GetSendEmailKey(emailStr, "")
	code, e := redisClient.Get(rkey).Result()
	if code == "" || e != nil {
		err = fmt.Errorf("验证码无效")
		return
	}
	if code != vcode {
		err = fmt.Errorf("验证码无效")
		return
	}
	db := s.baseService.Mysql.Get()
	var (
		userDto user.User
	)
	userFull.Roles = append(userFull.Roles, "admin")
	userDto, err = user.NewModel(db).ExistEmail(emailStr)
	authServer := auth_server.GetClient()
	if err != nil { //没有注册
		userDto.Email = emailStr
		userDto.Name = fmt.Sprintf("user_%d", time.Now().UnixNano())
		userDto, err = user.NewModel(db).Create(userDto)
		if err != nil {
			return
		}
		userFull.Uid = int64(userDto.Id)
		userFull.Name = userDto.Name
		userFull.Avatar = userDto.Avatar
		userFull.Email = userDto.Email
		userFull.CreatedAt = userDto.CreatedAt.Format("2006-01-02 15:04:05")
		userFull.LastLoginAt = userDto.LastLoginAt.Format("2006-01-02 15:04:05")
		req := &auth_server.CreateTokenReq{}
		req.TokenType = auth_server.TokenType_ADMIN
		req.Uid = int64(userDto.Id)
		ctokenRes, e := authServer.CreateToken(ctx, req)
		if e != nil {
			err = e
			return
		}
		if ctokenRes.Code != 200 {
			err = fmt.Errorf(ctokenRes.Msg)
			return
		}
		userFull.Token = ctokenRes.Token
		redisClient.Del(rkey)
		return
	}
	//注册了，自动登录吧
	userFull.Uid = int64(userDto.Id)
	userFull.Name = userDto.Name
	userFull.Avatar = userDto.Avatar
	userFull.Email = userDto.Email
	userFull.CreatedAt = userDto.CreatedAt.Format("2006-01-02 15:04:05")
	userFull.LastLoginAt = userDto.LastLoginAt.Format("2006-01-02 15:04:05")
	req := &auth_server.CreateTokenReq{}
	req.TokenType = auth_server.TokenType_ADMIN
	req.Uid = int64(userDto.Id)
	ctokenRes, e := authServer.CreateToken(ctx, req)
	if e != nil {
		err = e
		return
	}
	if ctokenRes.Code != 200 {
		err = fmt.Errorf(ctokenRes.Msg)
		return
	}
	userFull.Token = ctokenRes.Token
	redisClient.Del(rkey)
	_ = user.NewModel(db).UpdateLastLoginAtWithId(userDto.Id)
	return
}

func (s *UserService) DataFull(in user.User) (err error) {
	db := s.baseService.Mysql.Get()
	values := make(map[string]interface{})
	values["name"] = in.Name
	values["avatar"] = in.Avatar
	values["password"] = in.Password
	return user.NewModel(db).UpdateWithId(in.Id, values)
}

func (s *UserService) Login(emailStr, password string) (userFull *proto.RegisterRes, err error) {
	db := s.baseService.Mysql.Get()
	userDto, e := user.NewModel(db).FindWithLogin(emailStr, password)
	if e != nil {
		err = e
		return
	}
	userFull = new(proto.RegisterRes)
	userFull.Uid = int64(userDto.Id)
	userFull.Name = userDto.Name
	userFull.Avatar = userDto.Avatar
	userFull.Email = userDto.Email
	userFull.CreatedAt = userDto.CreatedAt.Format("2006-01-02 15:04:05")
	userFull.LastLoginAt = userDto.LastLoginAt.Format("2006-01-02 15:04:05")
	authServer := auth_server.GetClient()
	req := &auth_server.CreateTokenReq{}
	req.TokenType = auth_server.TokenType_ADMIN
	req.Uid = int64(userDto.Id)
	ctokenRes, e := authServer.CreateToken(context.Background(), req)
	if e != nil {
		err = e
		return
	}
	if ctokenRes.Code != 200 {
		err = fmt.Errorf(ctokenRes.Msg)
		return
	}
	userFull.Token = ctokenRes.Token
	userFull.Roles = append(userFull.Roles, "admin")
	_ = user.NewModel(db).UpdateLastLoginAtWithId(int(userFull.Uid))
	return
}

//
func (s *UserService) PasswordBack(emailStr, vcode, password string) (err error) {
	db := s.baseService.Mysql.Get()
	redisClient := s.baseService.Redis.Get()
	rkey := s.Redis.GetSendEmailKey(emailStr, "reset_password")
	code, e := redisClient.Get(rkey).Result()
	if code == "" || e != nil {
		err = fmt.Errorf("验证码无效")
		return
	}
	if code != vcode {
		err = fmt.Errorf("验证码无效")
		return
	}
	values := make(map[string]interface{})
	values["password"] = password
	return user.NewModel(db).UpdateWithEmail(emailStr, values)
}
