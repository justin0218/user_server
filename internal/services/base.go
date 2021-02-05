package services

import (
	"user_server/store"
)

type baseService struct {
	Mysql  store.Mysql
	Redis  store.Redis
	Config store.Config
}
