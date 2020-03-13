package user

import (
	"time"
)

const (
	Student = 0
	Teacher = 1
)

// 用户
type User struct {

	// 用户唯一标识
	ID string `json:"id" bson:"_id"`

	// 用户名
	Name string `json:"name" bson:"name"`

	// 微信接口返回的用户唯一标识,可能用不上，先保留
	OpenID string `json:"open_id" bson:"open_id"`

	// 用户手机号
	Phone string `json:"phone" bson:"phone"`

	// 邮箱
	Email string `json:"email" bson:"email"`

	// 密码
	Password string `json:"password" bson:"password"`

	// 用户类型: 暂定为:1教师 2学生
	UserType int `json:"user_type" bson:"user_type"`

	// 用户头像路径
	Avatar string `json:"avatar" bson:"avatar"`

	// 创建时间
	CreateTime time.Time `json:"create_time" bson:"create_time"`

	// 最后一次修改时间
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}

/*
type BaseDto struct {
	Name string `json:"name"`

	Password string `json:"password"`

	UserType int `json:"user_type"`
}

type WXLoginDto struct {
	BaseDto
	OpenID string `json:"open_id"`
}*/

// 检查user_type是否符合规范
func (u *User) CheckType() bool {
	return u.UserType == Teacher || u.UserType == Student
}
