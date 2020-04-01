package service

import (
	"github.com/NothingXiang/online-class/courseware"
)

type CoursewareService interface {
	CreateCourseware(c *courseware.CoursewareInfo) error

	RemoveCourseware(coursewareID string) error

	ListCourseware(classId string, page, limit int) ([]*courseware.CoursewareInfo, error)

	GetCourseware(wid string) (*courseware.CoursewareInfo, error)
}
