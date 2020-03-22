package service

import (
	"fmt"
	"log"
	"time"

	"github.com/NothingXiang/online-class/common/dbutil"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/notice"
	"github.com/NothingXiang/online-class/notice/store"
	jsoniter "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
)

const (
	// 通知存储在redis中的key, 占位符中填class id
	RedisNoticeKey = "Notice:%v"
)

type NoticeServiceImpl struct {
	store store.NoticeStore
}

func (n *NoticeServiceImpl) RemoveNotice(noticeID string) error {
	cid, err := n.store.RemoveNotice(noticeID)
	if err != nil {
		return err
	}
	//clear class notice cache
	dbutil.Redis().Del(fmt.Sprintf(RedisNoticeKey, cid))

	return nil
}

func (n *NoticeServiceImpl) GetNoticeByClass(classID string, page, limit int) ([]*notice.Notice, error) {

	// check from redis
	key := fmt.Sprintf(RedisNoticeKey, classID)
	field := fmt.Sprintf("%v:%v", page, limit)

	results, err := GetNoticeFromRedis(key, field)
	if err == nil {
		return results, nil
	}

	results, err = n.store.GetNoticeByClass(classID, (page-1)*limit, limit)
	if err != nil {
		return nil, resp.DBError.NewErr(err)
	}

	//set into redis
	dbutil.Redis().HSet(key, field, results)

	return results, nil
}

func (n *NoticeServiceImpl) CreateNotice(ntc *notice.Notice) error {

	ntc.ID = uuid.NewV4().String()
	ntc.CreateTime = time.Now()

	err := n.store.CreateNotice(ntc)
	if err != nil {
		return err
	}

	//clear class notice cache
	dbutil.Redis().Del(fmt.Sprintf(RedisNoticeKey, ntc.Class))

	return nil
}

func (n *NoticeServiceImpl) UpdateNotice(update *notice.Notice) error {

	update.UpdateTime = time.Now()

	err := n.store.UpdateNotice(update)
	if err != nil {
		return err
	}
	//clear class notice cache
	dbutil.Redis().Del(fmt.Sprintf(RedisNoticeKey, update.Class))

	return nil
}

func (n *NoticeServiceImpl) GetNoticeTemplate(noticeType string) ([]*notice.Template, error) {
	return n.store.GetNoticeTemplate(noticeType)
}

/*
func (n *NoticeServiceImpl) BatchCreateTemplate(tmpls []*notice.Template) error {
	return n.store.BatchCreateTemplate(tmpls)
}*/

// 从redis里的hash中获取通知列表
func GetNoticeFromRedis(key, field string) ([]*notice.Notice, error) {

	result := make([]*notice.Notice, 0)

	s, err := dbutil.Redis().HGet(key, field).Result()
	if err != nil {
		log.Printf("get notice %v from redis error:%v", key, err)
		return nil, err
	}

	err = jsoniter.UnmarshalFromString(s, &result)
	if err != nil {
		log.Printf("Unmarshal Notice %v from redis error:%v", s, err)
		return nil, err
	}

	return result, nil
}
