package services

import (
	"github.com/NothingXiang/online-class/user"
)

// service层返回的error必须带有具体的错误类型
type UserService interface {
	Create(user *user.User) error

	LoginByPwd(phone string, pwd string) (*user.User, error)

	//LoginById(id string, pwd string) (*user.User, error)

	//FindUser(id string) (*user.User, error)

	// 根据id和密码检查是否存在该用户
	CheckUserIdAndPwd(id, pwd string) error

	// 通过小程序回传的wechat code 来检查是否存在该用户
	CheckUserByWeChat(code string) (*user.User, error)
}
