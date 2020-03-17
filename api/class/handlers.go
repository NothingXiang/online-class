package class

import (
	"log"

	"github.com/NothingXiang/online-class/class"
	"github.com/NothingXiang/online-class/class/services"
	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/gin-gonic/gin"
)

var (
	cs services.ClassService
)

func init() {
	//	todo: 实例化ClassService...
}

func GetClassesByUser(c *gin.Context) {

	uid, ok := req.TryGetParam("uid", c)
	if !ok {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	cls, err := cs.GetClassByUser(uid)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(cls))
}

func CreateClass(c *gin.Context) {
	var cl class.Class

	if err := c.BindJSON(&cl); err != nil {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	// 1.参数校验
	if !req.CheckEmpty(c, cl.Name, cl.School, cl.MasterID) {
		return
	}

	// 2.入库
	err := cs.CreateClass(&cl)

	if err != nil {
		log.Println(err)
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(cl))

}

func ListTeacher(c *gin.Context) {
	cid, ok := req.TryGetParam("cid", c)
	if !ok {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	ts, err := cs.GetTeachers(cid)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(ts))
}

func ListStudentPageable(c *gin.Context) {
	// 1. 获取入参
	cid, ok := req.TryGetParam("cid", c)
	page := c.GetInt("page")
	limit := c.GetInt("limit")

	if !ok || !req.CheckPage(page, limit) {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	ts, err := cs.GetStudents(cid, page, limit)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(ts))
}
/*
func GetTeacher(c *gin.Context) {
	uid, ok := req.TryGetParam("user_id", c)
	if !ok {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	t, err := cs.GetTeacher(uid)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(t))
}*/

func GetStudent(c *gin.Context) {

}

func AddTeacher(c *gin.Context) {

	var t class.AddTeacherDto
	if err := c.BindJSON(&t); err != nil {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	err := cs.AddTeacher(&t)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(&t))

}

func AddStudent(c *gin.Context) {

	var t class.AddStudentDto
	if err := c.BindJSON(&t); err != nil {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	err := cs.AddStudent(&t)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(&t))
}

func DeleteSubject(c *gin.Context) {

	var cid, tid string
	var flag bool
	cid, flag = req.TryGetParam("cid", c)
	if !flag {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}
	tid, flag = req.TryGetParam("cid", c)
	if !flag {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	var s class.Subject
	s = class.Subject(c.GetInt("subject"))
	if !class.CheckSubject(s) {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	err := cs.RemoveSubject(cid, tid,s)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(nil))

}

func AddSubject(c *gin.Context) {

	var cid, tid string
	var flag bool
	var s class.Subject

	cid, flag = req.TryGetParam("cid", c)
	if !flag {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}
	tid, flag = req.TryGetParam("cid", c)
	if !flag {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	s = class.Subject(c.GetInt("subject"))

	if !class.CheckSubject(s) {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	err := cs.AddSubject(cid, tid, s)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(s))
}
