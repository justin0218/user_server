package api

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"user_server/api/proto"
	"user_server/internal/services"
	"user_server/store"
)

type UserSvr struct {
	AdminUserService services.AdminUserService
	ClientUserService services.ClientUserService
}

func GrpcServer() {
	conf := new(store.Config)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s", conf.Get().Etcd.Key))
	if err != nil {
		panic(err)
	}
	var opts []grpc.ServerOption
	svr := grpc.NewServer(opts...)
	proto.RegisterUserServer(svr, &UserSvr{})
	err = svr.Serve(lis)
	if err != nil {
		panic(err)
	}
}
