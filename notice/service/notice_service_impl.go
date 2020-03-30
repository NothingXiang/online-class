package service

import (
	"fmt"
	"log"
	"time"

	store3 "github.com/NothingXiang/online-class/class/store"
	"github.com/NothingXiang/online-class/common/dbutil"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/notice"
	"github.com/NothingXiang/online-class/notice/store"
	store2 "github.com/NothingXiang/online-class/user/store"
	jsoniter "github.com/json-iterator/go"
	uuid "github.com/satori/go.uuid"
)

const (
	// 通知存储在redis中的key, 占位符中填class id
	RedisNoticeKey = "Notice:%v"

	// 某条通知的已读列表 key为noticeID ，value是list， 存储用户id
	RedisReadNotice = "ReadNoticeList:%v"
)

type NoticeServiceImpl struct {
	NoticeStore store.NoticeStore
	UserStore   store2.UserStore
	ClassStore  store3.ClassStore
}

func (n *NoticeServiceImpl) GetReadList(noticeID string) ([]string, error) {
	return dbutil.Redis().SMembers(fmt.Sprintf(RedisReadNotice, noticeID)).Result()
}

func (n *NoticeServiceImpl) GetNotice(noticeID string) (*notice.Notice, error) {
	ntc, err := n.NoticeStore.GetNotice(noticeID)

	if err != nil {
		return nil, resp.DBError.NewErr(err)
	}

	return ntc, nil
}

func (n *NoticeServiceImpl) AddNoticeReadList(noticeID, userID string) error {

	dbutil.Redis().SAdd(fmt.Sprintf(RedisReadNotice, noticeID), userID)

	return nil
}

func (n *NoticeServiceImpl) RemoveNotice(noticeID string) error {
	cid, err := n.NoticeStore.RemoveNotice(noticeID)
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

	results, err = n.NoticeStore.GetNoticeByClass(classID, (page-1)*limit, limit)
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

	err := n.NoticeStore.CreateNotice(ntc)
	if err != nil {
		return err
	}

	//clear class notice cache
	dbutil.Redis().Del(fmt.Sprintf(RedisNoticeKey, ntc.Class))

	return nil
}

func (n *NoticeServiceImpl) UpdateNotice(update *notice.Notice) error {

	update.UpdateTime = time.Now()

	err := n.NoticeStore.UpdateNotice(update)
	if err != nil {
		return err
	}
	//clear class notice cache
	dbutil.Redis().Del(fmt.Sprintf(RedisNoticeKey, update.Class))

	return nil
}

func (n *NoticeServiceImpl) GetNoticeTemplate(noticeType string) ([]*notice.Template, error) {
	return n.NoticeStore.GetNoticeTemplate(noticeType)
}

/*
func (n *NoticeServiceImpl) BatchCreateTemplate(tmpls []*notice.Template) error {
	return n.NoticeStore.BatchCreateTemplate(tmpls)
}*/

// 从redis里的hash中获取通知列表
func GetNoticeFromRedis(key, field string) ([]*notice.Notice, error) {

	result := make([]*notice.Notice, 0)

	s, err := dbutil.Redis().HGet(key, field).Result()
	if err != nil {
		log.Printf("get notice %v from redis error:%v\n", key, err)
		return nil, err
	}

	err = jsoniter.UnmarshalFromString(s, &result)
	if err != nil {
		log.Printf("Unmarshal Notice %v from redis error:%v", s, err)
		return nil, err
	}

	return result, nil
}
