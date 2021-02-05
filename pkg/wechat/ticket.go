package wechat

import (
	"fmt"
	"github.com/parnurzeal/gorequest"
)

type Ticket struct {
	Errmsg  string `json:"errmsg"`
	Errcode int    `json:"errcode"`
	Ticket  string `json:"ticket"`
}

type TicketFromApi struct {
	Data struct {
		Ticket
	} `json:"data"`
}

func GetTicket() (ret TicketFromApi, err error) {
	rurl := fmt.Sprintf("http://momoman.cn/mall/v1/server/ticket")
	_, _, errs := gorequest.New().Get(rurl).EndStruct(&ret)
	if ret.Data.Errcode != 0 || len(errs) > 0 {
		err = fmt.Errorf("wechat get ticket err:%v code:%d msg:%s", errs, ret.Data.Errcode, ret.Data.Errmsg)
		return
	}
	return
}
