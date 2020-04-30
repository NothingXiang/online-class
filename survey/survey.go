package survey

import (
	"time"
)

// 问卷
type Survey struct {
	// 唯一标识
	ID string `json:"id" bson:"_id"`

	// 问卷所属班级
	ClassID string `json:"class_id" bson:"class_id"`

	// 问卷标题
	Title string `json:"title" bson:"title"`

	// 问卷描述
	Content string `json:"description" bson:"content"`

	// 题目
	Questions []*Question `json:"questions" bson:"questions"`

	// 创建时间
	CreateTime time.Time `json:"create_time" bson:"create_time"`

	//问卷截至时间
	EndTime time.Time `json:"end_time" bson:"end_time"`

	// 问卷创建人
	CreateBy string `json:"create_by" bson:"create_by"`
}

// 题目
type Question struct {

	// 题号，没啥意义
	//ID int `json:"id" bson:"_id"`

	// 所属的问卷
	//SurveyID string `json:"survey_id" bson:"survey_id"`

	// 问卷类型
	Type int `json:"type" bson:"type"`

	// 问题
	Ques string `json:"ques" bson:"ques"`

	// 选项
	Options []string `json:"options" bson:"options"`
}

const (
	// 题目类型：单选，多选，填空
	Single = iota
	Multi
	Filling
)

// 检查type是否符合
func CheckQuesType(t int) bool {
	return t == Single || t == Multi || t == Filling
}
