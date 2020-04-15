package req

import (
	"reflect"
	"testing"

	"github.com/spf13/viper"
)

func TestCodeToWeChat(t *testing.T) {

	viper.SetConfigFile("..\\..\\config\\config.json")
	viper.ReadInConfig()

	tests := []struct {
		name    string
		args    string
		want    *WeChatLoginResponse
		wantErr bool
	}{
		{
			"",
			"sdqwvfv139285",
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CodeToWeChat(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("CodeToWeChat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CodeToWeChat() got = %v, want %v", got, tt.want)
			}
		})
	}
}
