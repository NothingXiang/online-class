package homework

import (
	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/homework"
	"github.com/NothingXiang/online-class/homework/service"
	"github.com/gin-gonic/gin"
)

var (
	homeworkService service.HomeworkService
)

func init() {
	//	todo: 实例化homeworkService

	homeworkService = service.NewHomeworkServiceImpl()
}

func GetHomework(c *gin.Context) {
	workID, suc := req.TryGetParam("wid", c)

	if !suc {
		resp.Json(c, resp.ParamEmptyErr.NewErrStr("homework id empty"))
		return
	}

	homework, err := homeworkService.GetHomework(workID)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(homework))

}

func ListHomework(c *gin.Context) {
	classID, suc := req.TryGetParam("cid", c)
	if !suc {
		resp.Json(c, resp.ParamEmptyErr.NewErrStr("class id empty"))
		return
	}

	page, _ := req.TryGetInt("page", c)
	limit, _ := req.TryGetInt("limit", c)

	if !req.CheckPage(page, limit) {
		resp.Json(c, resp.InvalidParamErr)
		return
	}

	lists, err := homeworkService.GetHomeworkByClass(classID, page, limit)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(lists))

}

func CreateHomework(c *gin.Context) {

	var work homework.Homework

	if err := c.Bind(&work); err != nil {
		resp.Json(c, resp.InvalidParamErr)
		return
	}

	if !req.CheckEmpty(c, work.Title, work.Class, work.CreateBy) {
		return
	}

	err := homeworkService.CreateHomework(&work)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.SucJson(c, nil)
}

func UpdateHomework(c *gin.Context) {
	var upd homework.Homework

	err := c.Bind(&upd)
	if err != nil {
		resp.ErrJson(c, resp.InvalidParamErr)
		return
	}

	if !req.CheckEmpty(c, upd.ID) {
		return
	}

	err = homeworkService.UpdateHomework(&upd)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, nil)

}

func RemoveHomework(c *gin.Context) {
	workId, suc := req.TryGetParam("wid", c)
	if !suc {
		resp.ErrJson(c, resp.ParamEmptyErr)
		return
	}

	err := homeworkService.RemoveHomework(workId)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}
	resp.SucJson(c, nil)
}

func GetReadList(c *gin.Context) {
	workId, suc := req.TryGetParam("wid", c)
	if !suc {
		resp.ErrJson(c, resp.ParamEmptyErr)
		return
	}

	list, err := homeworkService.GetReadList(workId)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, list)
}

func AddReadList(c *gin.Context) {

	workID, suc := req.TryGetParam("wid", c)
	if !suc {
		resp.ErrJson(c, resp.ParamEmptyErr)
		return
	}
	userId, s := req.TryGetParam("uid", c)
	if !s {
		resp.ErrJson(c, resp.ParamEmptyErr)
		return
	}

	err := homeworkService.AddReadList(workID, userId)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, nil)
}
