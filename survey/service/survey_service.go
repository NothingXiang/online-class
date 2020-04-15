package service

import (
	"github.com/NothingXiang/online-class/survey"
)

// 问卷服务接口
type SurveyService interface {

	// 创建问卷
	CreateSurvey(s *survey.Survey) error

	// 通过id查找问卷
	GetSurveyByID(surveyID string) (*survey.Survey, error)

	// 分页获取班级问卷
	FindSurveyByClass(classID string, page, limit int) ([]*survey.Survey, error)

	// 删除某份问卷
	RemoveSurvey(surveyID string) error

	// 创建一份答卷
	CreateAnswer(a []*survey.AnswerSheet) error

	// 生成问卷统计信息
	GenerateStatistics(surveyID string) ([]*survey.Statistics, error)
}
