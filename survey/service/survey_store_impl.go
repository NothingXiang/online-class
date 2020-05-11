package service

import (
	"time"

	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/survey"
	"github.com/NothingXiang/online-class/survey/store"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type SurveyServiceImpl struct {
	SurveyStore store.SurveyStore
}

func (si *SurveyServiceImpl) CreateSurvey(s *survey.Survey) error {

	// 1.检查 问卷所属班级 和创建人的身份
	s.ID = uuid.NewV4().String()
	s.CreateTime=time.Now()

	// 2. 检查问题,不能为空
	if len(s.Questions) == 0 {
		return resp.ParamEmptyErr.NewErrStr(" survey question empty")
	}

	err := si.SurveyStore.CreateSurvey(s)
	if err != nil {
		logrus.Error(err)
		return resp.DBError.NewErr(err)

	}
	return nil

}

func (si *SurveyServiceImpl) GetSurveyByID(surveyID string) (*survey.Survey, error) {

	return si.SurveyStore.GetSurveyByID(surveyID)
}

func (si *SurveyServiceImpl) FindSurveyByClass(classID string, page, limit int) ([]*survey.Survey, error) {

	return si.SurveyStore.FindSurveyByClass(classID, (page-1)*limit, limit)
}

func (si *SurveyServiceImpl) RemoveSurvey(surveyID string) error {
	return si.SurveyStore.RemoveSurvey(surveyID)
}

func (si *SurveyServiceImpl) CreateAnswer(answers []*survey.AnswerSheet) error {

	if len(answers) == 0 {
		return resp.ParamEmptyErr
	}

	// 查找和答卷联系的那份问卷
	s, err := si.GetSurveyByID(answers[0].SurveyID)
	if err != nil {
		return err
	}

	// 比较问卷的截止时间
	//if time.Now().After(s.EndTime) {
	//	return resp.ForbiddenError.NewErrStr("survey is end")
	//}

	if len(s.Questions) != len(answers) {
		return resp.InvalidParamErr.NewErrStr("question and answer number match failed")
	}

	// 检查这一组答卷的合法性
	for index, answer := range answers {

		// 记录这一组题目可选的答案
		Optional := make(map[string]bool, len(s.Questions))
		for _, que := range s.Questions[index].Options {
			Optional[que] = true
		}

		//	1. 所属题号必须按顺序
		if answer.Sequence == index+1 {
			return resp.ParamFmtErr.NewErrStr("question number incorrect ")
		}

		// 2. 判断题型
		// todo: check answer match
		switch s.Questions[index].Type {
		case survey.Single:
			if len(answer.Selected) != 1 {
				return resp.InvalidParamErr.NewErrStr("answers can only select one")
			}

			if !Optional[answer.Selected[0]] {
				return resp.InvalidParamErr.NewErrStr("select a invalid option")
			}

		case survey.Multi:
			if len(answer.Selected) < 1 || len(answer.Selected) > len(s.Questions[index].Options) {
				return resp.InvalidParamErr.NewErrStr("answer select too much")
			}

			for _, ans := range answer.Selected {
				if !(Optional[ans]) {
					return resp.InvalidParamErr.NewErrStr("answer select too much")
				}
			}

		case survey.Filling:
			if len(answer.Selected) != 1 {
				return resp.ParamEmptyErr
			}
		default:
			logrus.Errorf("survey %v haves some error!", s.ID)
			return resp.UnknownError
		}

		answer.ID = uuid.NewV4().String()

	}
	// 入库
	return si.SurveyStore.CreateAnswer(answers)
}

//todo: unfinished
func (si *SurveyServiceImpl) GenerateStatistics(surveyID string) ([]*survey.Statistics, error) {

	sur, err := si.GetSurveyByID(surveyID)
	if err != nil {
		return nil, err
	}

	answers, e := si.GetAnswersBySurvey(surveyID)

	if e != nil {
		return nil, err
	}

	return Generate(sur, answers)

}

func (si *SurveyServiceImpl) GetAnswersBySurvey(surveyID string) ([]*survey.AnswerSheet, error) {

	return si.SurveyStore.GetAnswerBySurvey(surveyID)

}
