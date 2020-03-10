package dbutil

import (
	"log"

	"github.com/NothingXiang/online-class/config"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

func init() {

	config.Init()

	// 初始化mongo连接
	mgoUrl := viper.GetString("mgo.url")
	var err error
	if session, err = mgo.Dial(mgoUrl); err != nil {
		log.Fatalf("[NewSession] dial mongo failed:%v", err)
	}


	// 初始化redis连接
	Redis = redis.NewClient(&redis.Options{
		Addr:     config.GetDeStr("redis.addr", "localhost:6379"),
		Password: viper.GetString("redis.pwd"),
		DB:       viper.GetInt("redis.db"),
	})
	if _, err = Redis.Ping().Result(); err != nil {
		log.Fatalf("[NewRedisConnect] dial Redis failed:%v", err)
	}
}

func MgoDB() *mgo.Database {
	return session.Clone().DB(viper.GetString("mgo.db"))
}

// 返回对应的collection
func MongoColl(name string) *mgo.Collection {
	return MgoDB().C(name)
}
