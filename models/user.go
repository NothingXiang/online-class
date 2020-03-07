package models

import (
	"time"
)

type Users struct {

	// 用户唯一标识（感觉作用不大，用wx_code说不定更好）
	ID string `json:"id"`

	// 用户名
	Name string `json:"name"`

	// 微信接口返回的用户唯一标识,可能用不上，先保留
	OpenID string `json:"open_id"`

	// 用户手机号
	Phone string `json:"phone"`

	// 用户类型: 暂定为:1教师 2学生
	UserType int `json:"user_type"`

	// 用户头像路径
	Avatar string `json:"avatar"`

	// 创建时间
	CreateTime time.Time `json:"create_time"`

	// 最后一次修改时间
	UpdateTime time.Time `json:"update_time"`
}
