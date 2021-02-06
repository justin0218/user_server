package wechat_server

import (
	"sync"
	"user_server/pkg/etcd"
)

var once sync.Once
var conn WechatClient

func GetClient() WechatClient {
	once.Do(func() {
		conn = NewWechatClient(etcd.Discovery("wechat_server"))
	})
	return conn
}
