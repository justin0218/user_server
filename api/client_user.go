package api

import (
	"context"
	"user_server/api/proto"
)

func (s *UserSvr) ClientUserWechatLogin(ctx context.Context, req *proto.ClientUserWechatLoginReq) (ret *proto.ClientUserWechatLoginRes, err error) {
	ret = &proto.ClientUserWechatLoginRes{Code: 600}
	res,e := s.ClientUserService.Login(req.Code)
	if e != nil {
		ret.Msg = e.Error()
		return
	}
	ret = res
	ret.Code = 200
	return
}
