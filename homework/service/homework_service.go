package service

import (
	"github.com/NothingXiang/online-class/homework"
)

type HomeworkService interface {
	GetHomework(workId string) (*homework.Homework, error)

	GetReadList(workID string) ([]string, error)

	GetHomeworkByClass(classID string, page, limit int) ([]*homework.Homework, error)

	CreateHomework(work *homework.Homework) error

	RemoveHomework(workId string) error

	UpdateHomework(work *homework.Homework) error

	AddReadList(workID, userID string) error
}
