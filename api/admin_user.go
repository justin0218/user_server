package api

import (
	"context"
	"user_server/api/proto"
	"user_server/internal/models/admin_user"
)

func (s *UserSvr) AdminSendEmailCode(ctx context.Context, req *proto.AdminSendEmailCodeReq) (*proto.AdminSendEmailCodeRes, error) {
	ret := &proto.AdminSendEmailCodeRes{Code: 400}
	if req.Email == "" {
		ret.Msg = "邮箱发送失败，请检查邮箱是否正确"
		return ret, nil
	}
	e := s.AdminUserService.SendEmailCode(req.Email, req.From)
	if e != nil {
		ret.Msg = e.Error()
		return ret, nil
	}
	ret.Code = 200
	return ret, nil
}

func (s *UserSvr) AdminRegister(ctx context.Context, req *proto.AdminRegisterReq) (*proto.AdminRegisterRes, error) {
	ret := &proto.AdminRegisterRes{Code: 400}
	if req.Email == "" {
		ret.Msg = "邮箱发送失败，请检查邮箱是否正确"
		return ret, nil
	}
	r, e := s.AdminUserService.Register(req.Email, req.Code)
	if e != nil {
		ret.Msg = e.Error()
		return ret, nil
	}
	r.Code = 200
	return r, nil
}

func (s *UserSvr) AdminLogin(ctx context.Context, req *proto.AdminLoginReq) (*proto.AdminRegisterRes, error) {
	ret := &proto.AdminRegisterRes{Code: 400}
	if req.Email == "" {
		ret.Msg = "邮箱发送失败，请检查邮箱是否正确"
		return ret, nil
	}
	r, e := s.AdminUserService.Login(req.Email, req.Password)
	if e != nil {
		ret.Msg = e.Error()
		return ret, nil
	}
	r.Code = 200
	return r, nil
}

func (s *UserSvr) AdminPasswordBack(ctx context.Context, req *proto.AdminPasswordBackReq) (*proto.AdminPasswordBackRes, error) {
	ret := &proto.AdminPasswordBackRes{Code: 400}
	if req.Email == "" {
		ret.Msg = "请检查邮箱是否正确"
		return ret, nil
	}
	e := s.AdminUserService.PasswordBack(req.Email, req.Code, req.Password)
	if e != nil {
		ret.Msg = e.Error()
		return ret, nil
	}
	ret.Code = 200
	return ret, nil
}

func (s *UserSvr) AdminDataFull(ctx context.Context, req *proto.AdminDataFullReq) (*proto.Res, error) {
	ret := &proto.Res{Code: 400}
	if req.Uid <= 0 || req.Name == "" || req.Avatar == "" || req.Password == "" {
		ret.Msg = "参数错误"
		return ret, nil
	}
	e := s.AdminUserService.DataFull(admin_user.User{
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
