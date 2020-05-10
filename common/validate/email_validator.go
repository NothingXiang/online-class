package validate

import (
	"html/template"
	"strings"
	"time"

	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/common/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/gomail.v2"
)

var (
	Email = EmailValidator{}

	// 不加must,不然跑不了test
	EmailTmpl, _ = template.ParseFiles(".\\static\\template\\email.html")
)

// 邮箱验证
type EmailValidator struct {
}

func (e *EmailValidator) GenerateCode(toEmail string, expire time.Duration) error {

	// 1.check email format
	if !utils.IsEmailAddr(toEmail) {
		return resp.ParamFmtErr
	}

	// 2. generate random code
	code := GenerateCode(toEmail, 30*time.Minute)

	// 3. execute body
	var msg strings.Builder
	EmailTmpl.Execute(&msg, map[string]string{"code": code})

	// 4. send email
	err := SendMessage(toEmail, "【微课】验证码通知", &msg)
	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}

func SendMessage(to, subject string, body *strings.Builder) error {

	msg := gomail.NewMessage()

	msg.SetHeader("Subject", subject)
	msg.SetHeader("From", viper.GetString("email.account"))
	msg.SetHeader("To", to)

	msg.SetBody("text/html", body.String())

	err := gomail.NewDialer(
		viper.GetString("email.addr"),
		viper.GetInt("email.port"),
		viper.GetString("email.account"),
		viper.GetString("email.pwd"),
	).DialAndSend(msg)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
