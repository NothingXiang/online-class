package store

import (
	"github.com/NothingXiang/online-class/common/dbutil"
	"github.com/NothingXiang/online-class/courseware"
	"gopkg.in/mgo.v2/bson"
)

const (
	CoursewareClct = "courseware"
)

type CoursewareMgoStore struct {
}

func (cw *CoursewareMgoStore) CreateCourseware(c *courseware.CoursewareInfo) error {

	return dbutil.MongoColl(CoursewareClct).Insert(c)

}

func (cw *CoursewareMgoStore) RemoveCourseware(wareID string) error {

	return dbutil.MongoColl(CoursewareClct).RemoveId(wareID)
}

func (cw *CoursewareMgoStore) GetCourseware(wid string) (*courseware.CoursewareInfo, error) {

	var c courseware.CoursewareInfo

	err := dbutil.MongoColl(CoursewareClct).FindId(wid).One(&c)

	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (cw *CoursewareMgoStore) ListCourseware(classId string, skip, limit int) ([]*courseware.CoursewareInfo, error) {

	find := bson.M{
		"class": classId,
	}

	var cws []*courseware.CoursewareInfo

	err := dbutil.MongoColl(CoursewareClct).
		Find(find).
		Skip(skip).Limit(limit).
		Sort("-create_time").
		All(&cws)

	if err != nil {
		return nil, err
	}

	return cws, nil
}
