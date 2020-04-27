package store

import (
	"github.com/NothingXiang/online-class/class"
)

type ClassStore interface {

	// 创建班级
	CreateClass(class *class.Class) error

	//	通过id查找班级
	FindClassByID(cid string) (*class.Class, error)

	// 通过所有班级id符合的班级
	PatchClass(ids []string) ([]*class.Class, error)

	// 分页查找属于这个班级的学生
	FindStudent(classID string, skip, limit int) ([]*class.Student, error)

	// 查找用户所属的教师角色
	FindStudentByUser(userID string) ([]*class.Student, error)

	// 查找属于这个班级的教师
	FindTeacherByClass(classID string) ([]*class.Teacher, error)

	// 查找用户所属的教师角色
	FindTeacherByUser(userID string) ([]*class.Teacher, error)

	GetTeacher(userID, classID string) (*class.Teacher, error)

	//	通过学生id查找班级
	FindClassByStudent(sid string) (*class.Class, error)

	//	通过教师id查找班级
	FindClassByTeacher(tid string) (*class.Class, error)

	// 添加学生
	AddStudent(classID string, stu *class.Student) error

	// 添加班级
	AddTeacher(te *class.Teacher) error

	// 增加老师所教授的科目
	AddTechSubject(classID, teacherID string, s class.Subject) error

	// 减少老师所教授的科目
	DeleteTechSubject(classID, teacherID string, s class.Subject) error

	// 更新教师所教授的科目
	UpdateSubject(classID, teacherID string, s []class.Subject) error

	// 修改班主任
	UpdateMaster(classID, masterID string) error

	// 移除学生
	RemoveStudent(classID, studentID string) error

	// 移除教师
	RemoveTeacher(classID, teacherID string) error
}
