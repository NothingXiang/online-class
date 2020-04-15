package store

import (
	"github.com/NothingXiang/online-class/courseware"
)

type CoursewareStore interface {
	CreateCourseware(c *courseware.CoursewareInfo) error

	RemoveCourseware(wareID string) error

	GetCourseware(wid string) (*courseware.CoursewareInfo, error)

	ListCourseware(classId string, skip, limit int) ([]*courseware.CoursewareInfo, error)
}
