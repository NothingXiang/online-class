package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/user"
	"github.com/NothingXiang/online-class/user/store"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type UserServiceImpl struct {
	Store store.UserStore
}

func (u *UserServiceImpl) CreateByWeChat(dto *user.WeChatCrateDto) error {

	// 1. 用code向微信后台换取openid
	weChatData, err := req.CodeToWeChat(dto.Code)
	if err != nil {
		logrus.Error(err)
		return resp.WeChatError.NewErr(err)
	}

	// 2. 检查该账号是否已经存在
	exist, _ := u.Store.FindUserByOpenID(weChatData.OpenID)
	if exist != nil {
		return resp.RepeatError.NewErrStr(weChatData.OpenID)
	}

	// 3.创建账号
	dto.ID = uuid.NewV4().String()
	dto.OpenID = weChatData.OpenID
	dto.CreateTime = time.Now()
	err = u.Store.CreateUser(&dto.User)

	if err != nil {
		return resp.DBError.NewErr(err)
	}

	return nil
}

func (u *UserServiceImpl) CheckUserByWeChat(code string) (*user.User, error) {

	weChatData, err := req.CodeToWeChat(code)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	user, err := u.Store.FindUserByOpenID(weChatData.OpenID)

	if err != nil {
		logrus.Printf("find user failed:%v", err)
		return nil, err
	}

	return user, nil
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
