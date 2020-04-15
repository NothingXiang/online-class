package req

import (
	"encoding/json"

	"github.com/NothingXiang/online-class/common/resp"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// 微信登录接口
	WeChatCode2session = "https://api.weixin.qq.com/sns/jscode2session"
)

// 微信登录接口返回的响应数据
type WeChatLoginResponse struct {

	// 用户唯一标识
	OpenID string `json:"openid"`

	// 会话密钥
	SessionKey string `json:"sessionkey"`

	// 用户在开放平台的唯一标识符，在满足 UnionID 下发条件的情况下会返回，详见 UnionID 机制说明。
	UnionID string `json:"unionid"`

	// 错误码
	ErrCode int `json:"errcode"`

	// 错误信息
	ErrMsg string `json:"errmsg"`
}

// Example:
//GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
func CodeToWeChat(code string) (*WeChatLoginResponse, error) {

	var r WeChatLoginResponse

	cli := resty.New()
	response, err := cli.R().SetQueryParams(
		map[string]string{
			"appid":      viper.GetString("wechat.appid"),
			"secret":     viper.GetString("wechat.appsecret"),
			"js_code":    code,
			"grant_type": "authorization_code",
		}).Get(WeChatCode2session)

	if err != nil {
		logrus.Errorf("get WeChat jscode2session failed:%v", err)
		return nil, resp.OutDoorError.NewErr(err)
	}

	json.Unmarshal(response.Body(), &r)

	if r.ErrCode != 0 {
		logrus.Errorf("WeChat jscode2session return failed:%v", r.ErrMsg)
		return nil, resp.OutDoorError.NewErrStr(r.ErrMsg)
	}

	return &r, nil

}
