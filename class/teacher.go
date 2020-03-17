package class

// 教授科目
type Subject int

const (
	Chinese Subject = iota
	Maths
	English
	Science
	Physics
	Chemistry
	Biology
	Politics
	Moral
	History
	Geography
	Natural
	Sports
	Technology
	Art
	Music
	Others
)

type Teacher struct {

	// 唯一标识
	ID string `json:"id" bson:"id"`

	//外键，教师应该是用户中的一员,该id应该与user中某一位相同
	UserID string `json:"user_id" bson:"user_id"`

	// 外键，老师所属的班级
	ClassID string `json:"class_id" bson:"class_id"`

	// 老师所教授的科目
	Subjects []Subject `json:"subjects" bson:"subjects"`
}

// 校验subjects参数
func CheckSubject(t Subject) bool {

	return SubjectTable[t]
}

var SubjectTable = map[Subject]bool{
	Chinese:    true,
	Maths:      true,
	English:    true,
	Science:    true,
	Physics:    true,
	Chemistry:  true,
	Biology:    true,
	Politics:   true,
	Moral:      true,
	History:    true,
	Geography:  true,
	Natural:    true,
	Sports:     true,
	Technology: true,
	Art:        true,
	Music:      true,
	Others:     true,
}

type AddTeacherDto struct {
	Teacher
	ClassID string `json:"class_id"`
}
