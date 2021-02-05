package wechat

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type AccessToken struct {
	Errmsg      string `json:"errmsg"`
	Errcode     int    `json:"errcode"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type AccessTokenFromApi struct {
	Data struct {
		Errmsg      string `json:"errmsg"`
		Errcode     int    `json:"errcode"`
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	} `json:"data"`
}

func GetAccessToken() (ret AccessTokenFromApi, err error) {
	rurl := fmt.Sprintf("http://momoman.cn/mall/v1/server/access_token")
	_, _, errs := gorequest.New().Get(rurl).EndStruct(&ret)
	if ret.Data.Errcode != 0 || len(errs) > 0 {
		err = fmt.Errorf("wechat get access_token err:%v code:%d msg:%s", errs, ret.Data.Errcode, ret.Data.Errmsg)
		return
	}
	return
}
