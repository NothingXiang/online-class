package class

type Student struct {
	// 唯一标识
	ID string `json:"id" bson:"_id"`

	// 学生名字
	Name string `json:"name" bson:"name"`

	//	 外键,用户id
	UserID string `json:"user_id" bson:"user_id"`

	//	外键，所属班级
	ClassID string `json:"class_id" bson:"class_id"`
}
/*
type AddStudentDto struct {
	Student
	Class string `json:"class_id"`
}*/
