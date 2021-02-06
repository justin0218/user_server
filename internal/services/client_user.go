package services

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"user_server/api/auth_server"
	"user_server/api/proto"
	"user_server/api/wechat_server"
	"user_server/internal/models/client_user"
)

type ClientUserService struct {
	baseService
}

func (s *ClientUserService) Login(code string) (ret *proto.ClientUserWechatLoginRes, err error) {
	ret = new(proto.ClientUserWechatLoginRes)
	wechatServer := wechat_server.GetClient()
	ctx := context.Background()
	authToken, e := wechatServer.GetAuthAccessToken(ctx, &wechat_server.GetAuthAccessTokenReq{
		Account: wechat_server.Account_momo_za_huo_pu,
		Code:    code,
	})
	if e != nil {
		err = e
		return
	}
	if authToken.Res.Code != 200 {
		err = fmt.Errorf("%s", authToken.Res.Msg)
		return
	}
	wuserInfo, e := wechatServer.GetUserInfo(ctx, &wechat_server.GetUserInfoReq{
		AuthAccessToken: authToken.AccessToken,
		Openid:          authToken.Openid,
	})
	if e != nil {
		err = e
		return
	}
	if wuserInfo.Res.Code != 200 {
		err = fmt.Errorf("%s", wuserInfo.Res.Msg)
		return
	}
	db := s.Mysql.Get()
	authServer := auth_server.GetClient()

	olduser, e := client_user.NewModel(db).GetByOpenid(authToken.Openid)
	if e == gorm.ErrRecordNotFound { //未注册
		newuser, e := client_user.NewModel(db).Create(client_user.User{
			Openid:   wuserInfo.Openid,
			Avatar:   wuserInfo.Headimgurl,
			Nickname: wuserInfo.Nickname,
		})
		if e != nil {
			err = e
			return
		}
		ret.Openid = newuser.Openid
		ret.Nickname = newuser.Nickname
		ret.Avatar = newuser.Avatar
		ret.Uid = int64(newuser.Id)
		tokenRet, e := authServer.CreateToken(ctx, &auth_server.CreateTokenReq{
			Uid:       int64(newuser.Id),
			TokenType: auth_server.TokenType_CLIENT,
		})
		if e != nil {
			err = e
			return
		}
		if tokenRet.Code != 200 {
			err = fmt.Errorf("%s", tokenRet.Msg)
			return
		}
		ret.Token = tokenRet.Token
		return
	}
	if e != nil {
		err = e
		return
	}
	err = client_user.NewModel(db).UpdateById(client_user.User{
		Id:       olduser.Id,
		Avatar:   wuserInfo.Headimgurl,
		Nickname: wuserInfo.Nickname,
	})
	if err != nil {
		return
	}
	ret.Openid = olduser.Openid
	ret.Nickname = olduser.Nickname
	ret.Avatar = olduser.Avatar
	ret.Uid = int64(olduser.Id)
	tokenRet, e := authServer.CreateToken(ctx, &auth_server.CreateTokenReq{
		Uid:       int64(olduser.Id),
		TokenType: auth_server.TokenType_CLIENT,
	})
	if e != nil {
		err = e
		return
	}
	if tokenRet.Code != 200 {
		err = fmt.Errorf("%s", tokenRet.Msg)
		return
	}
	ret.Token = tokenRet.Token
	return
}

func (s *ClientUserService) ClientGetUserByUid(uid int64) (ret *proto.ClientGetUserByUidRes, err error) {
	ret = new(proto.ClientGetUserByUidRes)
	clientUser := client_user.NewModel(s.Mysql.Get())
	uinfo, e := clientUser.GetByUid(int(uid))
	if e != nil {
		err = e
		return
	}
	ret.Nickname = uinfo.Nickname
	ret.Openid = uinfo.Openid
	ret.Avatar = uinfo.Avatar
	ret.Uid = uid
	return
}
