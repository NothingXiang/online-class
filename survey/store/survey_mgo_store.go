package store

import (
	"github.com/NothingXiang/online-class/common/dbutil"
	"github.com/NothingXiang/online-class/survey"
	"gopkg.in/mgo.v2/bson"
)

const (
	SurveyClct = "survey"
	AnswerClct = "answer"
)

type SurveyMgoStore struct {
}

func (sm *SurveyMgoStore) GetAnswerBySurvey(surveyID string) ([]*survey.AnswerSheet, error) {

	var answers []*survey.AnswerSheet

	err := dbutil.MongoColl(AnswerClct).
		Find(bson.M{"survey_id": surveyID}).
		All(&answers)

	if err != nil {
		return nil, err
	}

	return answers, nil

}

func (sm *SurveyMgoStore) CreateSurvey(s *survey.Survey) error {

	return dbutil.MongoColl(SurveyClct).Insert(s)
}

func (sm *SurveyMgoStore) GetSurveyByID(surveyID string) (*survey.Survey, error) {

	var s survey.Survey

	err := dbutil.MongoColl(SurveyClct).FindId(surveyID).One(&s)

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (sm *SurveyMgoStore) GetSurveyByUser(userID string, skip, limit int) ([]*survey.Survey, error) {

	var ss []*survey.Survey

	err := dbutil.MongoColl(SurveyClct).
		Find(bson.M{"create_by": userID}).
		Skip(skip).Limit(limit).Sort("-create_time").
		All(&ss)

	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (sm *SurveyMgoStore) FindSurveyByClass(classId string, skip, limit int) ([]*survey.Survey, error) {

	var ss []*survey.Survey

	err := dbutil.MongoColl(SurveyClct).
		Find(bson.M{"class_id": classId}).
		Skip(skip).Limit(limit).Sort("-create_time").
		All(&ss)

	if err != nil {
		return nil, err
	}
	return ss, nil
}

func (sm *SurveyMgoStore) RemoveSurvey(surveyID string) error {

	return dbutil.MongoColl(SurveyClct).RemoveId(surveyID)
}

func (sm *SurveyMgoStore) CreateAnswer(a []*survey.AnswerSheet) error {

	return dbutil.MongoColl(AnswerClct).Insert(a)
}
