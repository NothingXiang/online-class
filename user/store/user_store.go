package store

import (
	"github.com/NothingXiang/online-class/user"
)

// user存储层的相关接口
type UserStore interface {

	// 创建
	CreateUser(user *user.User) error

	// 删除
	DeleteUser(id string) error

	//	改
	UpdateUser(user *user.User) error

	// 查
	FindUser(id string) (*user.User, error)

	// 通过微信openID查找用户
	FindUserByOpenID(openID string) (*user.User, error)

	//
	FindUserByPwd(name, pwd string) (*user.User, error)

	// 检查手机号是否重复
	IsPhoneRepect(phone string) bool

	FindUserByIdandPwd(id, pwd string) (*user.User, error)
}
