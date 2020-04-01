package service

import (
	"github.com/NothingXiang/online-class/courseware"
	"github.com/NothingXiang/online-class/courseware/store"
	uuid "github.com/satori/go.uuid"
)

type CoursewareServiceImpl struct {
	CoursewareStore store.CoursewareStore
}

func (cs *CoursewareServiceImpl) CreateCourseware(c *courseware.CoursewareInfo) error {

	c.ID = uuid.NewV4().String()

	return cs.CoursewareStore.CreateCourseware(c)

}

func (cs *CoursewareServiceImpl) RemoveCourseware(coursewareID string) error {

	return cs.CoursewareStore.RemoveCourseware(coursewareID)
}

func (cs *CoursewareServiceImpl) ListCourseware(classId string, page, limit int) ([]*courseware.CoursewareInfo, error) {

	return cs.CoursewareStore.ListCourseware(classId, (page-1)*limit, limit)
}

func (cs *CoursewareServiceImpl) GetCourseware(wid string) (*courseware.CoursewareInfo, error) {
	return cs.CoursewareStore.GetCourseware(wid)
}
