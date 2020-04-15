package service

import (
	"fmt"
	"time"

	"github.com/NothingXiang/online-class/class"
	"github.com/NothingXiang/online-class/common/dbutil"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/homework"
	"github.com/NothingXiang/online-class/homework/store"
	uuid "github.com/satori/go.uuid"
)

const (
	RedisHomeworkKey = "homework:%v"

	RedisReadHomework = "ReadHomeworkList:%v"
)

type HomeworkServiceImpl struct {
	HomeworkStore store.HomeworkStore
}

func NewHomeworkServiceImpl() *HomeworkServiceImpl {
	return &HomeworkServiceImpl{HomeworkStore: &store.HomeworkMgoStore{}}
}

func (h *HomeworkServiceImpl) GetHomework(workId string) (*homework.Homework, error) {

	work, err := h.HomeworkStore.GetHomework(workId)
	if err != nil {
		return nil, resp.DBError.NewErr(err)
	}

	return work, nil
}

func (h *HomeworkServiceImpl) GetReadList(workID string) ([]string, error) {

	return dbutil.Redis().SMembers(fmt.Sprintf(RedisReadHomework, workID)).Result()
}

func (h *HomeworkServiceImpl) GetHomeworkByClass(classID string, page, limit int) (res []*homework.Homework, err error) {

	// check from redis
	key := fmt.Sprintf(RedisHomeworkKey, classID)
	field := fmt.Sprintf("%v:%v", page, limit)
	err = dbutil.GetHashFromRedis(key, field, &res)
	if err == nil {
		return res, nil
	}

	// if get from redis failed,check from database
	res, err = h.HomeworkStore.GetHomeworkByClass(classID, (page-1)*limit, limit, "-create_time")
	if err != nil {
		return nil, resp.DBError.NewErr(err)
	}

	// add to redis
	dbutil.Redis().HSet(key, field, res)

	return res, nil
}

func (h *HomeworkServiceImpl) CreateHomework(work *homework.Homework) error {

	if !class.CheckSubject(work.Subject) {
		return resp.InvalidParamErr.NewErrStr(fmt.Sprintf("subject %v", work.Subject))
	}

	work.ID = uuid.NewV4().String()
	work.CreateTime = time.Now()

	err := h.HomeworkStore.CreateHomework(work)

	if err != nil {
		return resp.DBError.NewErr(err)
	}

	return nil
}

func (h *HomeworkServiceImpl) RemoveHomework(workId string) error {
	cid, err := h.HomeworkStore.RemoveHomework(workId)

	if err != nil {
		return err
	}
	dbutil.Redis().Del(fmt.Sprintf(RedisHomeworkKey, cid))

	return nil
}

func (h *HomeworkServiceImpl) UpdateHomework(work *homework.Homework) error {
	work.UpdateTime = time.Now()

	err := h.HomeworkStore.UpdateHomework(work)
	if err != nil {
		return err
	}
	dbutil.Redis().Del(fmt.Sprintf(RedisHomeworkKey, work.Class))

	return nil
}

func (h *HomeworkServiceImpl) AddReadList(workID, userID string) error {
	dbutil.Redis().SAdd(fmt.Sprintf(RedisReadHomework, workID), userID)
	return nil
}
