package services

import (
	"github.com/NothingXiang/online-class/class"
)

type ClassService interface {

	// 创建班级
	CreateClass(class *class.Class) error

	// 根据用户id和类型获取班级
	GetClassByUser(uid string, userType int) ([]*class.Class, error)

	// 通过用户id获取教师
	//GetTeacher(userID string) (class.Teacher, error)

	// 通过班级id获取教师列表
	GetTeachers(cid string) ([]*class.Teacher, error)

	// 分页获取学生列表
	GetStudents(cid string, page, limit int) ([]*class.Student, error)

	// 班级增加学生
	AddStudent(dto *class.Student) error

	// 班级开除学生
	RemoveStudent(classID, studentID string) error

	// 增加教师
	AddTeacher(dto *class.AddTeacherDto) error

	// 班级开除教师
	RemoveTeacher(classID, teacherID string) error

	// 减少教师在该班级担任的科目
	RemoveSubject(classID string, TeacherID string, s class.Subject) error

	AddSubject(classID string, TeacherID string, subject class.Subject) error

	UpdateSubject(classID string, TeacherID string, subject []class.Subject) error

	// 通过用户id和班级id获取教师信息
	GetTeacher(userID string, ClassID string) (*class.Teacher, error)

	// 通过用户id和班级id获取教师信息
	GetStudent(userID string, ClassID string) (*class.Student, error)
}
