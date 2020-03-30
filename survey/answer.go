package survey

// 问卷回答
type AnswerSheet struct {
	ID string `json:"id" bson:"_id"`

	// 所属用户
	UserID string `json:"user_id" bson:"user_id"`

	// 所属问卷
	SurveyID string `json:"survey_id"`

	// 所属题号
	Sequence int `json:"Sequence" bson:"Sequence"`

	// 所选的选项
	Selected []string `json:"selected" bson:"selected"`
}

// 对某一个问题的统计结果
type Statistics struct {

	// 题目
	Title string `json:"title"`

	// 题目类型
	Type int `json:"type"`

	// 选项
	Options []string `json:"options"`

	// 每个选项对应选择的人数
	Count map[string]int `json:"count"`
}
