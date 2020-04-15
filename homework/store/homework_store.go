package store

import (
	"github.com/NothingXiang/online-class/homework"
)

type HomeworkStore interface {

	// 获取某项作业
	GetHomework(workID string) (*homework.Homework, error)

	// 分页获取班级作业
	GetHomeworkByClass(classId string, skip, limit int, order string) ([]*homework.Homework, error)

	// 发布作业
	CreateHomework(hw *homework.Homework) error

	// 删除作业
	RemoveHomework(workId string) (classID string, err error)

	// 编辑作业
	UpdateHomework(update *homework.Homework) error


}
