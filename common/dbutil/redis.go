package dbutil

import (
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	log "github.com/sirupsen/logrus"
)

var (
	redisClient *redis.Client
)

// 从redis里的hash中获取通知列表
func GetHashFromRedis(key, field string, result interface{}) error {

	res, err := redisClient.HGet(key, field).Result()

	if err != nil {
		log.Printf(" get hash {%v:%v} result from redis err:%v\n", key, field, err)
		return err
	}

	err = jsoniter.UnmarshalFromString(res, &result)

	if err != nil {
		log.Printf("unmarshal data error:%v\n", err)
		return err
	}

	return nil

}
