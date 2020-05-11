package class

import (
	"encoding/json"
	"fmt"
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
	cs = services.NewClassServiceImpl()
}

func GetClassById(c *gin.Context) {
	classId, suc := req.TryGetParam("cid", c)
	if !suc {
		resp.ErrJson(c, resp.ParamEmptyErr)
		return
	}

	classData, err := cs.GetClassById(classId)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, classData)
}

func GetClassesByUser(c *gin.Context) {

	uid, ok := req.TryGetParam("uid", c)
	userType, ok2 := req.TryGetInt("type", c)
	if !ok || !ok2 {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	cls, err := cs.GetClassByUser(uid, userType)

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
	page, _ := req.TryGetInt("page", c)
	limit, _ := req.TryGetInt("limit", c)

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

func GetTeacher(c *gin.Context) {
	uid, ok := req.TryGetParam("uid", c)
	if !ok {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	cid, ok2 := req.TryGetParam("cid", c)
	if !ok2 {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	t, err := cs.GetTeacher(uid, cid)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(t))
}

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

	var t class.Student
	if err := c.BindJSON(&t); err != nil {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	if !req.CheckEmpty(c, t.ClassID, t.UserID, t.Name) {
		return
	}

	err := cs.AddStudent(&t)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(&t))
}

// todo: 可以改成传个dto...一个个获取参数也太累了.....
func DeleteSubject(c *gin.Context) {

	var cid, tid string
	var flag bool
	cid, flag = req.TryGetParam("cid", c)
	if !flag {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}
	tid, flag = req.TryGetParam("tid", c)
	if !flag {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	s, _ := req.TryGetInt("subject", c)
	if !class.CheckSubject(class.Subject(s)) {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	err := cs.RemoveSubject(cid, tid, class.Subject(s))
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(nil))

}

func AddSubject(c *gin.Context) {

	var cid, tid string
	var flag bool

	cid, flag = req.TryGetParam("cid", c)
	if !flag {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}
	tid, flag = req.TryGetParam("tid", c)
	if !flag {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	s, _ := req.TryGetInt("subject", c)

	if !class.CheckSubject(class.Subject(s)) {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	err := cs.AddSubject(cid, tid, class.Subject(s))

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(s))
}

func UpdateSubject(c *gin.Context) {

	var cid, tid string
	var flag bool

	//cid, flag = req.TryGetParam("cid", c)
	//if !flag {
	//	resp.Json(c, resp.ParamEmptyErr)
	//	return
	//}
	tid, flag = req.TryGetParam("tid", c)
	if !flag {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	res := c.Query("subjects")
	fmt.Println(res)

	var sub []class.Subject

	if err := json.Unmarshal([]byte(res), &sub); err != nil {
		resp.ErrJson(c, resp.ParamFmtErr)
		return
	}

	for _, s := range sub {
		if !class.CheckSubject(s) {
			resp.Json(c, resp.ParamFmtErr)
			return
		}
	}

	err := cs.UpdateSubject(cid, tid, sub)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, nil)

}

type subs struct {
	Subjects []class.Subject
}
