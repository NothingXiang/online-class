package courseware

import (
	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/courseware"
	"github.com/NothingXiang/online-class/courseware/service"
	"github.com/gin-gonic/gin"
)

var (
	coursewareService service.CoursewareService
)

func init() {
	//	todo： init service
	coursewareService = service.NewCoursewareServiceImpl()
}

func CreateCourseware(c *gin.Context) {

	//  get info
	var cw courseware.CoursewareInfo
	if err := c.Bind(&cw); err != nil {
		resp.ErrJson(c, resp.InvalidParamErr)
		return
	}

	// check info
	if !req.CheckEmpty(c, cw.Name, cw.CreateBy, cw.Class) {
		return
	}

	// save info
	err := coursewareService.CreateCourseware(&cw)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, nil)

}

func RemoveCourseware(c *gin.Context) {
	coursewareID, suc := req.TryGetParam("cid", c)

	if !suc {
		resp.ErrJson(c, resp.ParamEmptyErr)
		return
	}

	err := coursewareService.RemoveCourseware(coursewareID)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, nil)

}

func ListCoursewareInfo(c *gin.Context) {

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

	cs, err := coursewareService.ListCourseware(classID, page, limit)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, cs)
}

func GetCoursewareInfo(c *gin.Context) {

	wid, suc := req.TryGetParam("wid", c)

	if !suc {
		resp.ErrJson(c, resp.ParamEmptyErr)
		return
	}

	courseware, err := coursewareService.GetCourseware(wid)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, courseware)

}
