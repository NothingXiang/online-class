package validate

import (
	"html/template"
	"testing"
	"time"

	"github.com/spf13/viper"
)

func TestEmailValidator_GenerateCode(t *testing.T) {

	// 因为测试文件的工作目录跟main的工作目录不同，引入外部文件的地方都必须重新加载
	EmailTmpl = template.Must(template.ParseFiles("..\\..\\static\\template\\email.html"))
	viper.SetConfigFile("..\\..\\config\\config.json")
	viper.ReadInConfig()

	type args struct {
		toEmail string
		expire  time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"",
			args{
				// todo : change to you test email address
				toEmail: "wei_class@foxmail.com",
				expire:  30 * time.Minute,
			},
			false,
		},
	}
	for _, tt := range tests {
		//for i := 0; i < 10; i++ {
		t.Run(tt.name, func(t *testing.T) {
			e := &EmailValidator{}
			if err := e.GenerateCode(tt.args.toEmail, tt.args.expire); (err != nil) != tt.wantErr {
				t.Errorf("GenerateCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		//}

	}
}
