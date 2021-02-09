package api

import (
	"context"
	"user_server/api/proto"
)

func (s *UserSvr) ClientCreateUser(ctx context.Context, req *proto.ClientCreateUserReq) (ret *proto.ClientCreateUserRes, err error) {
	ret = new(proto.ClientCreateUserRes)
	ret, err = s.ClientUserService.ClientCreateUser(req)
	if err != nil {
		return
	}
	return
}

func (s *UserSvr) ClientUpdateByUid(ctx context.Context, req *proto.ClientUpdateByUidReq) (ret *proto.ClientUpdateByUidRes, err error) {
	ret = new(proto.ClientUpdateByUidRes)
	ret, err = s.ClientUserService.UpdateById(req)
	if err != nil {
		return
	}
	return
}

func (s *UserSvr) ClientGetUserByOpenid(ctx context.Context, req *proto.ClientGetUserByOpenidReq) (ret *proto.ClientGetUserByOpenidRes, err error) {
	ret = new(proto.ClientGetUserByOpenidRes)
	ret, err = s.ClientUserService.GetUserByOpenid(req.Openid)
	if err != nil {
		return
	}
	return
}

func (s *UserSvr) ClientGetUserByUid(ctx context.Context, req *proto.ClientGetUserByUidReq) (ret *proto.ClientGetUserByUidRes, err error) {
	ret = &proto.ClientGetUserByUidRes{Code: 600}
	res, e := s.ClientUserService.ClientGetUserByUid(req.Uid)
	if e != nil {
		ret.Msg = e.Error()
		return
	}
	ret = res
	ret.Code = 200
	return
}
