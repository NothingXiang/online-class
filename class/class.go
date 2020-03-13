package class

// 班级
type Class struct {

	// 班级id
	ID string `json:"id" bson:"_id"`

	// 班级名称
	Name string `json:"name" bson:"name"`

	// 所属学校
	School string `json:"school" bson:"school"`

	// 班主任id
	MasterID string `json:"master_id" bson:"master_id"`

	//
	Teachers []string `json:"teachers" bson:"teachers"`

	//	学生
	students []string `json:"students"`
}
