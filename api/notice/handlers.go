package notice

import (
	store3 "github.com/NothingXiang/online-class/class/store"
	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/notice"
	"github.com/NothingXiang/online-class/notice/service"
	"github.com/NothingXiang/online-class/notice/store"
	store2 "github.com/NothingXiang/online-class/user/store"
	"github.com/gin-gonic/gin"
)

var (
	ns service.NoticeService
)

func init() {

	ns = &service.NoticeServiceImpl{
		NoticeStore: &store.NoticeMgoStore{},
		UserStore:   &store2.UserMgoStore{},
		ClassStore:  &store3.ClassMgoStore{},
	}
}

func UpdateNotice(c *gin.Context) {
	var ntc notice.Notice

	if err := c.Bind(&ntc); err != nil {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	if !req.CheckEmpty(c, ntc.ID) {
		return
	}

	if err := ns.UpdateNotice(&ntc); err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(ntc))
}

func GetTemplate(c *gin.Context) {
	typ, ok := req.TryGetParam("type", c)
	if !ok {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	templates, err := ns.GetNoticeTemplate(typ)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(templates))

}

func CreateNotice(c *gin.Context) {
	var ntc notice.Notice

	if err := c.Bind(&ntc); err != nil {
		resp.Json(c, resp.ParamFmtErr)
		return
	}

	if !req.CheckEmpty(c, ntc.Title, ntc.CreateBy, ntc.Class) {
		return
	}

	if err := ns.CreateNotice(&ntc); err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(ntc))

}

func GetNoticeByClassPageable(c *gin.Context) {
	cid, ok := req.TryGetParam("cid", c)
	page, _ := req.TryGetInt("page", c)
	limit, _ := req.TryGetInt("limit", c)

	if !ok || !req.CheckPage(page, limit) {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	class, err := ns.GetNoticeByClass(cid, page, limit)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(class))
}

// 添加某条通知的已读列表
func AddNoticeRead(c *gin.Context) {

	nid, ok := req.TryGetParam("nid", c)

	uid, ok2 := req.TryGetParam("uid", c)

	if !ok || !ok2 {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	ns.AddNoticeReadList(nid, uid)

	resp.Json(c, resp.NewSucResp(nil))

	return

}

// 获取某个通知的已读列表
func GetReadList(c *gin.Context) {

	nid, ok := req.TryGetParam("nid", c)
	if !ok {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	list, err := ns.GetReadList(nid)
	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(list))
}
