package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/user"
	"github.com/NothingXiang/online-class/user/store"
	uuid "github.com/satori/go.uuid"
)

type UserServiceImpl struct {
	Store store.UserStore
}

func (u *UserServiceImpl) CheckUserIdAndPwd(id, pwd string) error {

	user, e := u.Store.FindUserByIdandPwd(id, pwd)
	if e != nil {
		return resp.UnknownError.NewErr(e)
	}
	if user == nil {
		return resp.NotExistError
	}

	return nil
}

func (u *UserServiceImpl) LoginByPwd(phone string, pwd string) (*user.User, error) {
	user, err := u.Store.FindUserByPwd(phone, pwd)

	if err != nil {
		return nil, resp.NotExistError.NewErr(err)
	}

	return user, nil
}

//
//func (u *UserServiceImpl) LoginById(id string, pwd string) (*user.User, error) {
//	panic("implement me")
//}

func (u *UserServiceImpl) Create(user *user.User) error {
	// check phone repeat

	if u.Store.IsPhoneRepect(user.Phone) {
		return resp.RepeatError.NewErr(errors.New("phone number:" + user.Phone))
	}

	// 生成新的uuid
	user.ID = uuid.NewV4().String()

	user.CreateTime = time.Now()

	if err := u.Store.CreateUser(user); err != nil {
		return fmt.Errorf("[CreateUserByPhonePwd] failed:%w", err)
	}

	return nil
}

/*
func (u *UserServiceImpl) Login() error {
	panic("implement me")
}

func (u *UserServiceImpl) Logout() error {
	panic("implement me")
}

func (u *UserServiceImpl) FindUser(id string) (*user.User, error) {
	panic("implement me")
}
*/
