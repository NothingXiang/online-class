package services

import (
	"github.com/NothingXiang/online-class/class"
	"github.com/NothingXiang/online-class/class/store"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/user"
	store2 "github.com/NothingXiang/online-class/user/store"
	uuid "github.com/satori/go.uuid"
)

type ClassServiceImpl struct {
	cs store.ClassStore
	us store2.UserStore
}

// 创建一个班级，只有班主任才能创建班级
func (c *ClassServiceImpl) CreateClass(cl *class.Class) error {
	// 1. 检查创建人的身份,必须是教师
	// todo：可以单独拉出来作为权限系统
	master, err := c.us.FindUser(cl.MasterID)
	if err != nil {
		return resp.DBError.NewErr(err)
	}

	if master.UserType != user.Teacher {
		return resp.NotAuthError.NewErrStr("user type unmatched")
	}

	// 2. 补充部分参数
	cl.ID = uuid.NewV4().String()

	teach := class.Teacher{
		ID:      uuid.NewV4().String(),
		UserID:  master.ID,
		ClassID: cl.ID,
	}

	// 创建班级
	if e := c.cs.CreateClass(cl); e != nil {
		return resp.DBError.NewErr(e)
	}

	// 增加教师
	if e := c.cs.AddTeacher(&teach); e != nil {
		return resp.DBError.NewErr(e)
	}

	return nil
}

func (c *ClassServiceImpl) GetClassByUser(uid string, userType int) ([]*class.Class, error) {

	ids := make([]string, 0)

	switch userType {
	case user.Teacher:
		ts, _ := c.cs.FindTeacherByUser(uid)
		for _, t := range ts {
			ids = append(ids, t.ClassID)
		}
	case user.Student:
		ss, _ := c.cs.FindStudentByUser(uid)
		for _, s := range ss {
			ids = append(ids, s.ClassID)

		}
	}

	return c.cs.PatchClass(ids)

}

/*func (c *ClassServiceImpl) GetTeacher(userID string) (class.Teacher, error) {
	// 1. 获取所有

	panic("implement me")
}
*/
func (c *ClassServiceImpl) GetTeachers(cid string) (ts []*class.Teacher, err error) {
	ts, err = c.cs.FindTeacherByClass(cid)

	return
}

func (c *ClassServiceImpl) GetStudents(cid string, page, limit int) ([]*class.Student, error) {

	//todo: 可以把一次全部查询保存在缓存中，以后每次从缓存中读取
	ss, err := c.cs.FindStudent(cid, (page-1)*limit, limit)

	return ss, err
}

func (c *ClassServiceImpl) AddStudent(dto *class.AddStudentDto) error {

	return c.cs.AddStudent(dto.ClassID, &dto.Student)
}

func (c *ClassServiceImpl) RemoveStudent(classID, studentID string) error {
	return c.cs.RemoveStudent(classID, studentID)
}

func (c *ClassServiceImpl) AddTeacher(dto *class.AddTeacherDto) error {
	return c.cs.AddTeacher(&dto.Teacher)
}

func (c *ClassServiceImpl) RemoveTeacher(classID, teacherID string) error {
	return c.cs.RemoveTeacher(classID, teacherID)
}

func (c *ClassServiceImpl) RemoveSubject(classID string, TeacherID string, s class.Subject) error {

	return c.cs.DeleteTechSubject(classID, TeacherID, s)
}

func (c *ClassServiceImpl) AddSubject(classID string, TeacherID string, subject class.Subject) error {
	return c.cs.AddTechSubject(classID, TeacherID, subject)
}
