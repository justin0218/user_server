package email

import (
	"gopkg.in/gomail.v2"
)

func Send(email, content, subject string) (err error) {
	m := gomail.NewMessage()
	m.SetHeader("From", "jiketang <13720009841@163.com>")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	d := gomail.NewDialer("smtp.163.com", 465, "13720009841@163.com", "930218hx")
	if err = d.DialAndSend(m); err != nil {
		return
	}
	return
}
