package survey

import (
	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/survey"
	"github.com/NothingXiang/online-class/survey/service"
	"github.com/NothingXiang/online-class/survey/store"
	"github.com/gin-gonic/gin"
)

var surveyService service.SurveyService

func init() {
	surveyService = &service.SurveyServiceImpl{SurveyStore: &store.SurveyMgoStore{}}
}

func GetStatistics(c *gin.Context) {
	sid, ok := req.TryGetParam("sid", c)
	if !ok {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}
	if !req.CheckEmpty(c, sid) {
		return
	}

	s, err := surveyService.GenerateStatistics(sid)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(s))

}

func CreateAnswer(c *gin.Context) {
	var ans []*survey.AnswerSheet

	if err := c.Bind(&ans); err != nil {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	err := surveyService.CreateAnswer(ans)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(ans))

}

func DeleteSurvey(c *gin.Context) {
	sid, ok := req.TryGetParam("sid", c)

	if !ok || sid == "" {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	err := surveyService.RemoveSurvey(sid)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(nil))
}

func ListSurvey(c *gin.Context) {

	classID, _ := req.TryGetParam("cid", c)

	page, _ := req.TryGetInt("page", c)

	limit, _ := req.TryGetInt("limit", c)

	if !req.CheckPage(page, limit) {
		resp.Json(c, resp.InvalidParamErr)
		return
	}

	surveys, err := surveyService.FindSurveyByClass(classID, page, limit)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(surveys))

}

func GetSurvey(c *gin.Context) {
	sid, ok := req.TryGetParam("sid", c)
	if !ok {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	sur, err := surveyService.GetSurveyByID(sid)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(sur))
}

func CreateSurvey(c *gin.Context) {
	var s survey.Survey

	if err := c.Bind(&s); err != nil {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	if !req.CheckEmpty(c, s.ClassID, s.Title, s.CreateBy) {
		return
	}

	err := surveyService.CreateSurvey(&s)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(s))
}
