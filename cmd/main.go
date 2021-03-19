package main

import (
	"fmt"
	"time"
	"user_server/api"
	"user_server/store"
)

func main() {
	redis := new(store.Redis)
	err := redis.Get().Ping().Err()
	if err != nil {
		panic(err)
	}
	mysql := new(store.Mysql)
	mysql.Get()
	log := new(store.Log)
	//go etcd.Register()
	go api.GrpcServer()
	log.Get().Debug("server started at %v", time.Now())
	fmt.Printf("server started at %v \n", time.Now())
	select {}
}
