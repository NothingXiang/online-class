package store

import (
	"github.com/NothingXiang/online-class/common/dbutil"
	"github.com/NothingXiang/online-class/user"
	"gopkg.in/mgo.v2/bson"
)

const (
	UserClct = "user"
)

type UserMgoStore struct {
}

func (u *UserMgoStore) FindUserByOpenID(openID string) (*user.User, error) {
	find := bson.M{
		"open_id": openID,
	}
	var user user.User
	err := dbutil.MongoColl(UserClct).Find(find).One(&user)

	if err != nil  {
		return nil, err
	}

	return &user, nil
}

func (u *UserMgoStore) FindUserByIdandPwd(id, pwd string) (*user.User, error) {
	find := bson.M{
		"_id":      id,
		"password": pwd,
	}
	var user user.User
	if err := dbutil.MongoColl(UserClct).Find(find).One(&user); err != nil {
		return nil, err
	}

	return &user, nil

}

func (u *UserMgoStore) FindUserByPwd(name, pwd string) (*user.User, error) {
	panic("implement me")
}

func (u *UserMgoStore) IsPhoneRepect(phone string) bool {

	find := bson.M{
		"phone": phone,
	}

	num, _ := dbutil.MongoColl(UserClct).Find(find).Count()

	return num != 0
}

func (u *UserMgoStore) CreateUser(user *user.User) error {

	return dbutil.MgoDB().C(UserClct).Insert(user)
}

func (u *UserMgoStore) DeleteUser(id string) error {
	return dbutil.MgoDB().C(UserClct).RemoveId(id)
}

func (u *UserMgoStore) UpdateUser(user *user.User) error {

	upd := bson.M{
		"name":  user.Name,
		"phone": user.Phone,
	}

	if err := dbutil.MongoColl(UserClct).UpdateId(user.ID, upd); err != nil {
		return err
	}

	return nil

}

func (u *UserMgoStore) FindUser(id string) (*user.User, error) {
	var user user.User

	if err := dbutil.MongoColl(UserClct).FindId(id).One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
