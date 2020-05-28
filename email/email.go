package email

import (
	"github.com/FengWuTech/commons/setting"
	"github.com/go-gomail/gomail"
)

func SendMail(title string, body string, to []string) {
	m := gomail.NewMessage()
	m.SetHeader("From", setting.EmailSetting.User)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(setting.EmailSetting.Host, setting.EmailSetting.Port, setting.EmailSetting.User, setting.EmailSetting.Password)

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
