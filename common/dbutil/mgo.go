package dbutil

import (
	"log"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"online-class/config"
)

var (
	session *mgo.Session
)

func init() {

	config.Init()

	mgoUrl := viper.GetString("mgo.url")
	var err error
	if session, err = mgo.Dial(mgoUrl); err != nil {
		log.Fatalf("[NewSession] dial mongo failed:%v", err)
	}
}

func MgoDB() *mgo.Database {
	return session.Clone().DB(viper.GetString("mgo.db"))
}
