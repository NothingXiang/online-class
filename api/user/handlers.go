package user

import (
	"log"
	"net/http"

	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/common/utils"
	"github.com/NothingXiang/online-class/user"
	"github.com/NothingXiang/online-class/user/services"
	"github.com/NothingXiang/online-class/user/store"
	"github.com/gin-gonic/gin"
)

var (
	us services.UserService
)

func init() {
	us = &services.UserServiceImpl{
		Store: &store.UserMgoStore{},
	}
}

//
func LoginByWeChatCode(c *gin.Context) {
	code, ok := req.TryGetParam("code", c)
	if !ok {
		resp.Json(c, resp.ParamEmptyErr)
		return
	}

	user, err := us.CheckUserByWeChat(code)

	if err != nil {
		resp.Json(c, resp.ErrResp(err))
		return
	}

	resp.Json(c, resp.NewSucResp(user))

}

// create account by wechat data
func CreateByWeChat(c *gin.Context) {

	// 1. get param
	var dto user.WeChatCrateDto
	if err := c.Bind(&dto); err != nil {
		resp.ErrJson(c, resp.ParamFmtErr)
		return
	}

	// 2. check param
	if !req.CheckEmpty(c, dto.Code, dto.Avatar) {
		return
	}
	if !dto.CheckType() {
		resp.ErrJson(c, resp.InvalidParamErr.NewErrStr("user_type"))
		return
	}

	err := us.CreateByWeChat(&dto)

	if err != nil {
		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, dto.User)

}

// 通过手机号和密码登录
func LoginByPhonePwd(c *gin.Context) {
	var user user.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusOK, resp.ParamFmtErr)
		return
	}

	// 参数校验
	if !req.CheckEmpty(c, user.Phone, user.Password) {
		return
	}
	if !utils.IsPhoneNum(user.Phone) {
		c.JSON(http.StatusOK, resp.ErrResp(resp.ParamFmtErr))
		return
	}

	me, err := us.LoginByPwd(user.Phone, user.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, resp.ErrResp(err))
		return
	}

	c.JSON(http.StatusOK, resp.NewSucResp(me))

}

// 通过手机号和密码来创建
func CreateUserByPhonePwd(c *gin.Context) {

	var user user.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusOK, resp.ParamFmtErr.NewErr(err))
		return
	}

	// 参数校验
	if !req.CheckEmpty(c, user.Name, user.Phone, user.Password) {
		return
	}

	// todo：测试模式下，不检验这个
	/*if !utils.IsPhoneNum(user.Phone) {
		c.JSON(http.StatusOK, resp.InvalidParamErr.NewErr(errors.New("phone number")))
		return
	}*/
	if !user.CheckType() {
		c.JSON(http.StatusOK, resp.InvalidParamErr)
		return
	}

	e := us.Create(&user)
	if e != nil {
		log.Println(e)
		c.JSON(http.StatusOK, resp.ErrResp(e))
		return
	}

	c.JSON(http.StatusOK, resp.NewSucResp(user))
}

func UpdateUser(c *gin.Context) {
	var u user.User

	if err := c.BindJSON(&u); err != nil {
		resp.Json(c, resp.InvalidParamErr)
		return
	}

	if !req.CheckEmpty(c, u.ID) {
		return
	}

	if err := us.UpdateUser(&u); err != nil {

		resp.ErrJson(c, err)
		return
	}

	resp.SucJson(c, nil)
}
