package api

import (
	"user_server/api/proto"
	"user_server/internal/models/user"
	"context"
)

func (s *AdminUserSvr) SendEmailCode(ctx context.Context, req *proto.SendEmailCodeReq) (*proto.SendEmailCodeRes, error) {
	ret := &proto.SendEmailCodeRes{Code: 400}
	if req.Email == "" {
		ret.Msg = "邮箱发送失败，请检查邮箱是否正确"
		return ret, nil
	}
	e := s.userService.SendEmailCode(req.Email, req.From)
	if e != nil {
		ret.Msg = e.Error()
		return ret, nil
	}
	ret.Code = 200
	return ret, nil
}

func (s *AdminUserSvr) Register(ctx context.Context, req *proto.RegisterReq) (*proto.RegisterRes, error) {
	ret := &proto.RegisterRes{Code: 400}
	if req.Email == "" {
		ret.Msg = "邮箱发送失败，请检查邮箱是否正确"
		return ret, nil
	}
	r, e := s.userService.Register(req.Email, req.Code)
	if e != nil {
		ret.Msg = e.Error()
		return ret, nil
	}
	r.Code = 200
	return r, nil
}

func (s *AdminUserSvr) Login(ctx context.Context, req *proto.LoginReq) (*proto.RegisterRes, error) {
	ret := &proto.RegisterRes{Code: 400}
	if req.Email == "" {
		ret.Msg = "邮箱发送失败，请检查邮箱是否正确"
		return ret, nil
	}
	r, e := s.userService.Login(req.Email, req.Password)
	if e != nil {
		ret.Msg = e.Error()
		return ret, nil
	}
	r.Code = 200
	return r, nil
}

func (s *AdminUserSvr) PasswordBack(ctx context.Context, req *proto.PasswordBackReq) (*proto.PasswordBackRes, error) {
	ret := &proto.PasswordBackRes{Code: 400}
	if req.Email == "" {
		ret.Msg = "请检查邮箱是否正确"
		return ret, nil
	}
	e := s.userService.PasswordBack(req.Email, req.Code, req.Password)
	if e != nil {
		ret.Msg = e.Error()
		return ret, nil
	}
	ret.Code = 200
	return ret, nil
}

func (s *AdminUserSvr) DataFull(ctx context.Context, req *proto.DataFullReq) (*proto.Res, error) {
	ret := &proto.Res{Code: 400}
	if req.Uid <= 0 || req.Name == "" || req.Avatar == "" || req.Password == "" {
		ret.Msg = "参数错误"
		return ret, nil
	}
	e := s.userService.DataFull(user.User{
		Id:       int(req.Uid),
		Name:     req.Name,
		Avatar:   req.Avatar,
		Password: req.Password,
	})
	if e != nil {
		ret.Msg = e.Error()
		return ret, nil
	}
	ret.Code = 200
	return ret, nil
}
