package services

import (
	"github.com/jinzhu/gorm"
	"user_server/api/proto"
	"user_server/internal/models/client_user"
)

type ClientUserService struct {
	baseService
}

func (s *ClientUserService) UpdateById(req *proto.ClientUpdateByUidReq) (ret *proto.ClientUpdateByUidRes, err error) {
	ret = new(proto.ClientUpdateByUidRes)
	db := s.Mysql.Get()
	err = client_user.NewModel(db).UpdateById(client_user.User{
		Id:       req.Uid,
		Avatar:   req.Avatar,
		Nickname: req.Nickname,
	})
	if err != nil {
		return
	}
	return
}

func (s *ClientUserService) GetUserByOpenid(openid string) (ret *proto.ClientGetUserByOpenidRes, err error) {
	ret = new(proto.ClientGetUserByOpenidRes)
	db := s.Mysql.Get()
	oldUser, e := client_user.NewModel(db).GetByOpenid(openid)
	if e == gorm.ErrRecordNotFound {
		ret.Code = 404 //未注册
		return
	}
	if e != nil {
		err = e
		return
	}
	ret.Openid = oldUser.Openid
	ret.Nickname = oldUser.Nickname
	ret.Avatar = oldUser.Avatar
	ret.Uid = int64(oldUser.Id)
	return
}

func (s *ClientUserService) ClientCreateUser(req *proto.ClientCreateUserReq) (ret *proto.ClientCreateUserRes, err error) {
	ret = new(proto.ClientCreateUserRes)
	db := s.Mysql.Get()
	newUser, e := client_user.NewModel(db).Create(client_user.User{
		Openid:   req.Openid,
		Avatar:   req.Avatar,
		Nickname: req.Nickname,
	})
	if e != nil {
		err = e
		return
	}
	ret.Uid = int64(newUser.Id)
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
