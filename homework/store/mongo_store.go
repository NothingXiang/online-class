package store

import (
	"github.com/NothingXiang/online-class/common/dbutil"
	"github.com/NothingXiang/online-class/homework"
	"gopkg.in/mgo.v2/bson"
)

const (
	HomeworkClct = "homework"
)

type HomeworkMgoStore struct {
}

func (h *HomeworkMgoStore) GetHomework(workID string) (*homework.Homework, error) {

	var work homework.Homework

	err := dbutil.MongoColl(HomeworkClct).FindId(workID).One(&work)

	if err != nil {
		return nil, err
	}

	return &work, nil

}

func (h *HomeworkMgoStore) GetHomeworkByClass(classId string, skip, limit int, order string) ([]*homework.Homework, error) {

	find := bson.M{
		"class": classId,
	}

	var works []*homework.Homework

	err := dbutil.MongoColl(HomeworkClct).
		Find(find).Skip(skip).Limit(limit).
		Sort("-" + order).
		All(&works)

	if err != nil {
		return nil, err
	}

	return works, nil
}

func (h *HomeworkMgoStore) CreateHomework(hw *homework.Homework) error {

	return dbutil.MongoColl(HomeworkClct).Insert(hw)
}

func (h *HomeworkMgoStore) RemoveHomework(workId string) (classID string, err error) {

	var del homework.Homework
	dbutil.MongoColl(HomeworkClct).FindId(workId).Select(bson.M{"class": 1}).One(&del)

	err = dbutil.MongoColl(HomeworkClct).RemoveId(workId)
	if err != nil {
		return "", err
	}

	return del.Class, nil
}

func (h *HomeworkMgoStore) UpdateHomework(update *homework.Homework) error {

	updator := bson.M{
		"update_time": update.UpdateTime,
	}

	if update.Title != "" {
		updator["title"] = update.Title
	}

	if update.Content != "" {
		updator["content"] = update.Content
	}

	if len(update.Photos) != 0 {
		updator["photos"] = update.Photos
	}

	err := dbutil.MongoColl(HomeworkClct).UpdateId(update.ID, bson.M{"$set": updator})

	return err
}
