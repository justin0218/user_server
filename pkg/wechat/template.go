package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/parnurzeal/gorequest"
	"time"
)

type TemplateItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type Template struct {
	Touser     string                  `json:"touser"`
	TemplateId string                  `json:"template_id"`
	Url        string                  `json:"url"`
	Data       map[string]TemplateItem `json:"data"`
}

type TemplateRes struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func SendTemplate(data Template) (err error) {
	accessToken, e := GetAccessToken()
	if e != nil {
		err = e
		return
	}
	sendData, e := json.Marshal(data)
	if e != nil {
		err = e
		return
	}
	rurl := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken.Data.AccessToken)
	res := TemplateRes{}
	_, _, errs := gorequest.New().Post(rurl).Timeout(time.Second*30).Set("Content-Type", "application/json").Send(string(sendData)).EndStruct(&res)
	if len(errs) != 0 {
		return errs[0]
	}
	if res.Errcode != 0 {
		err = fmt.Errorf(res.Errmsg)
		return
	}
	return nil
}
