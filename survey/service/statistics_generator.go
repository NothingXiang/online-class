package service

import (
	"fmt"

	"github.com/NothingXiang/online-class/survey"
	"github.com/sirupsen/logrus"
)

// todo: 重点重构下这里
// 生成统计结果
func Generate(sur *survey.Survey, ans []*survey.AnswerSheet) ([]*survey.Statistics, error) {

	result := make([]*survey.Statistics, len(sur.Questions))

	// 一（问卷）对多（问题）对多（答案）对多（填写者） ，搞不来
	for _, answer := range ans {

		if result[answer.Sequence] == nil {
			result[answer.Sequence] =
				InitStatistics(sur.Questions[answer.Sequence])
		}

		// 根据回答的题号得到这题的题型，根据题型再做处理
		switch sur.Questions[answer.Sequence].Type {
		case survey.Single:
			if len(answer.Selected) != 1 {
				return nil, fmt.Errorf("answers can only select one:%v", answer)
			}

			// 计数加1
			result[answer.Sequence].Count[answer.Selected[0]]++
		case survey.Multi:
			if len(answer.Selected) < 1 ||
				len(answer.Selected) > len(sur.Questions[answer.Sequence].Options) {
				return nil, fmt.Errorf("multi question select number error:%v", answer)
			}

			for _, s := range answer.Selected {
				result[answer.Sequence].Count[s]++
			}

		case survey.Filling:
			if len(answer.Selected) != 1 {
				return nil, fmt.Errorf("filling question can only answer once:%v", answer)
			}
			// 计数加1
			result[answer.Sequence].Count[answer.Selected[0]]++

		default:
			logrus.Error("some wrong question struct into this ")
			return nil, fmt.Errorf("some wrong question struct into this")
		}

	}

	return result, nil
}

// 从问题类中获取统计所需要的基本信息
func InitStatistics(question *survey.Question) *survey.Statistics {

	// 预先估计一下要分配的空间
	var length int
	switch question.Type {
	case survey.Single:
		length = 1
	case survey.Multi:
		length = len(question.Options)
	default:
		length = 0
	}

	return &survey.Statistics{
		Title:   question.Ques,
		Type:    question.Type,
		Options: question.Options,
		Count:   make(map[string]int, length),
	}
}
