package store

import (
	"github.com/NothingXiang/online-class/survey"
)

type SurveyStore interface {
	CreateSurvey(s *survey.Survey) error

	GetSurveyByID(surveyID string) (*survey.Survey, error)

	// 分页获取某个用户的问卷
	GetSurveyByUser(userID string, skip, limit int) ([]*survey.Survey, error)

	FindSurveyByClass(classId string, skip, limit int) ([]*survey.Survey, error)

	// 删除某份问卷
	RemoveSurvey(surveyID string) error

	// 创建一份答卷
	CreateAnswer(a []*survey.AnswerSheet) error

	// 获取一份问卷下所有答卷
	GetAnswerBySurvey(surveyID string) ([]*survey.AnswerSheet, error)

	// 生成问卷统计信息
	//GenerateStatistics(surveyID string) ([]*survey.Statistics, error)
}
