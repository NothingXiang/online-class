package store

import (
	"github.com/NothingXiang/online-class/class"
	"github.com/NothingXiang/online-class/common/dbutil"
	"gopkg.in/mgo.v2/bson"
)

const (
	ClassClct   = "class"
	TeacherClct = "teacher"
	StudentClct = "student"
)

type ClassMgoStore struct {
}

func (c *ClassMgoStore) CreateClass(class *class.Class) error {

	return dbutil.MgoDB().C(ClassClct).Insert(class)
}

func (c *ClassMgoStore) FindClassByID(cid string) (cl *class.Class, err error) {
	err = dbutil.MongoColl(ClassClct).FindId(cid).One(&cl)

	if err != nil {
		return nil, err
	}

	return cl, nil
}

func (c *ClassMgoStore) PatchClass(ids []string) (cs []*class.Class, err error) {
	find := bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}

	err = dbutil.MongoColl(ClassClct).Find(find).All(&cs)

	if err != nil {
		return nil, err
	}

	return cs, nil
}

func (c *ClassMgoStore) FindStudent(classID string, skip, limit int) ([]*class.Student, error) {
	find := bson.M{
		"class_id": classID,
	}

	var s []*class.Student

	err := dbutil.MongoColl(StudentClct).Find(find).Skip(skip).Limit(limit).All(&s)

	if err != nil {
		return nil, err
	}

	return s, nil

}

func (c *ClassMgoStore) FindStudentByUser(userID string) ([]*class.Student, error) {
	find := bson.M{
		"user_id": userID,
	}

	var s []*class.Student

	err := dbutil.MongoColl(StudentClct).Find(find).All(&s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (c *ClassMgoStore) FindTeacherByClass(classID string) ([]*class.Teacher, error) {
	find := bson.M{
		"class_id": classID,
	}

	var s []*class.Teacher

	err := dbutil.MongoColl(TeacherClct).Find(find).All(&s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (c *ClassMgoStore) FindTeacherByUser(userID string) ([]*class.Teacher, error) {
	find := bson.M{
		"user_id": userID,
	}

	var s []*class.Teacher

	err := dbutil.MongoColl(TeacherClct).Find(find).All(&s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (c *ClassMgoStore) FindClassByStudent(sid string) (*class.Class, error) {

	var stu class.Student

	err := dbutil.MongoColl(StudentClct).FindId(sid).One(&stu)

	if err != nil {
		return nil, err
	}

	return c.FindClassByID(stu.ClassID)
}

func (c *ClassMgoStore) FindClassByTeacher(tid string) (*class.Class, error) {
	var tea class.Teacher

	err := dbutil.MongoColl(TeacherClct).FindId(tid).One(&tea)

	if err != nil {
		return nil, err
	}

	return c.FindClassByID(tea.ClassID)
}

func (c *ClassMgoStore) AddStudent(classID string, stu *class.Student) error {

	return dbutil.MongoColl(StudentClct).Insert(stu)
}

func (c *ClassMgoStore) AddTeacher(te *class.Teacher) error {
	return dbutil.MongoColl(TeacherClct).Insert(te)
}

func (c *ClassMgoStore) AddTechSubject(classID, teacherID string, s class.Subject) error {

	selector := bson.M{
		"_id": teacherID,
	}

	update := bson.M{
		"$push": bson.M{
			"subjects": s,
		},
	}

	return dbutil.MongoColl(TeacherClct).Update(selector, update)

}

func (c *ClassMgoStore) DeleteTechSubject(classID, teacherID string, s class.Subject) error {

	selector := bson.M{
		"_id": teacherID,
	}

	update := bson.M{
		"$pull": bson.M{
			"subjects": s,
		},
	}

	return dbutil.MongoColl(TeacherClct).Update(selector, update)
}

func (c *ClassMgoStore) UpdateMaster(classID, masterID string) error {

	update := bson.M{
		"$set": bson.M{"master_id": masterID},
	}
	return dbutil.MongoColl(ClassClct).UpdateId(classID, update)
}

func (c *ClassMgoStore) RemoveStudent(classID, studentID string) error {

	return dbutil.MongoColl(StudentClct).RemoveId(studentID)
}

func (c *ClassMgoStore) RemoveTeacher(classID, teacherID string) error {
	return dbutil.MongoColl(TeacherClct).RemoveId(teacherID)
}
