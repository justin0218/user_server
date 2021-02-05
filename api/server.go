package api

import (
	"user_server/api/proto"
	"user_server/internal/services"
	"user_server/store"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type UserSvr struct {
	userService services.UserService
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
