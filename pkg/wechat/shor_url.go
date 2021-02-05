package wechat

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type ShorUrl struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	ShortUrl string `json:"short_url"`
}

func GetShortUrl(lurl string) (ret ShorUrl, err error) {
	accessToken, e := GetAccessToken()
	if e != nil {
		err = e
		return
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/shorturl?access_token=%s", accessToken.Data.AccessToken)
	_, _, errs := gorequest.New().Post(url).SendString(fmt.Sprintf(`{"action":"long2short","long_url":"%s"}`, lurl)).EndStruct(&ret)
	if len(errs) > 0 {
		err = errs[0]
		return
	}
	if ret.Errcode != 0 {
		err = fmt.Errorf(ret.Errmsg)
		return
	}
	return
}
