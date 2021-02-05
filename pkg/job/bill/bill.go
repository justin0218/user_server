package bill

import (
	"user_server/internal/services"
	"user_server/pkg/wechat"
	"user_server/store"
	"github.com/robfig/cron"
	"time"
)

func Run() {
	config := new(store.Config)
	authService := new(services.AuthService)
	c := cron.New()
	spec := "0 0 18 * * ?"
	_ = c.AddFunc(spec, func() {
		users := []int{1, 4, 5}
		for _, uid := range users {
			uinfo, _ := authService.GetUserInfo(uid)
			data := make(map[string]wechat.TemplateItem)
			data["first"] = wechat.TemplateItem{Value: "您好，今天是您预约的记账提醒日，请记得记账！"}
			data["keyword1"] = wechat.TemplateItem{Value: time.Now().Format("2006-01-02 15:04:05")}
			data["keyword2"] = wechat.TemplateItem{Value: time.Now().Add(time.Hour * 24).Format("2006-01-02 15:04:05")}
			data["keyword3"] = wechat.TemplateItem{Value: "每天"}
			data["remark"] = wechat.TemplateItem{Value: "您好，今天是您预约的记账提醒日，请点击记账！"}
			_ = wechat.SendTemplate(wechat.Template{
				Touser:     uinfo.Openid,
				TemplateId: config.Get().Wechat.BillNoticeTemplate,
				Url:        config.Get().Wechat.BillNoticeUrl,
				Data:       data,
			})
		}
	})
	c.Start()
}
