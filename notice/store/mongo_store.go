package store

import (
	"github.com/NothingXiang/online-class/common/dbutil"
	"github.com/NothingXiang/online-class/notice"
	"gopkg.in/mgo.v2/bson"
)

const (
	NoticeClct     = "notice"
	NoticeTmplCLct = "notice_template"
)

type NoticeMgoStore struct {
}

func (n *NoticeMgoStore) GetNoticeByClass(classID string, skip, limit int) (ns []*notice.Notice, err error) {

	query := bson.M{
		"class": classID,
	}

	err = dbutil.MongoColl(NoticeClct).
		Find(query).
		Skip(skip).
		Limit(limit).
		Sort("-create_time").
		All(&ns)

	if err != nil {
		return nil, err
	}

	return ns, nil

}

func (n *NoticeMgoStore) CreateNotice(notice *notice.Notice) error {

	return dbutil.MongoColl(NoticeClct).Insert(notice)

}

func (n *NoticeMgoStore) RemoveNotice(noticeID string) (string, error) {

	var del notice.Notice
	dbutil.MongoColl(NoticeClct).Find(noticeID).Select(bson.M{"class": 1}).One(&del)

	err := dbutil.MongoColl(NoticeClct).RemoveId(noticeID)

	if err != nil {
		return "", err
	}

	return del.Class, nil

}

func (n *NoticeMgoStore) UpdateNotice(update *notice.Notice) error {
	updator := bson.M{
		"$set": bson.M{
			"title":       update.Title,
			"content":     update.Content,
			"update_time": update.UpdateTime,
			"photos":      update.Photos,
		},
	}

	err := dbutil.MongoColl(NoticeClct).UpdateId(update.ID, updator)

	return err
}

func (n *NoticeMgoStore) GetNoticeTemplate(noticeType string) (nts []*notice.Template, err error) {

	selector := bson.M{"tmpl_type": noticeType}

	err = dbutil.MongoColl(NoticeTmplCLct).Find(selector).All(&nts)

	if err != nil {
		return nil, err
	}

	return nts, nil

}
