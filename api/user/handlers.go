package user

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/common/utils"
	"github.com/NothingXiang/online-class/config"
	user2 "github.com/NothingXiang/online-class/user"
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
func GetAccountByWeChat(c *gin.Context) {
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

func LoginByWeChat(c *gin.Context) {

	//	 1. get mini-program code
	//code, _ := req.TryGetParam("code", c)

}

// 通过手机号和密码登录
func LoginByPhonePwd(c *gin.Context) {
	var user user2.User
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

	var user user2.User
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

	//	 todo:入库操作...
	e := us.Create(&user)
	if e != nil {
		log.Println(e)
		c.JSON(http.StatusOK, resp.ErrResp(e))
		return
	}

	c.JSON(http.StatusOK, resp.NewSucResp(user))
}

// 上传用户头像
func UploadAvatar(c *gin.Context) {

	id := c.PostForm("id")
	pwd := c.PostForm("pwd")
	if id == "" || pwd == "" {
		c.JSON(http.StatusOK,
			resp.ErrResp(resp.ParamEmptyErr))
		return
	}

	if err := us.CheckUserIdAndPwd(id, pwd); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, resp.ErrResp(err))
		return
	}

	file, err := c.FormFile("uploadAvatar")
	if err != nil {
		c.JSON(http.StatusOK, resp.UnknownError.NewErr(err))
		return
	}

	dir := fmt.Sprintf("%v/%v",
		config.GetDeStr("avatar.dir", "./avatar"),
		id)

	// 创建目录并且保存文件
	os.MkdirAll(dir, os.ModeDir)
	if e := c.SaveUploadedFile(file, dir+"/avatar.jpg"); e != nil {
		c.JSON(http.StatusOK,
			resp.UnknownError.NewErr(e))
		return
	}

	c.JSON(http.StatusOK, resp.NewSucResp(nil))

}

// 获取用户头像
func GetAvatar(c *gin.Context) {

	id := c.Query("id")
	pwd := c.Query("pwd")
	if id == "" || pwd == "" {
		c.JSON(http.StatusOK,
			resp.ErrResp(resp.ParamEmptyErr))
		return
	}

	if err := us.CheckUserIdAndPwd(id, pwd); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, resp.ErrResp(err))
		return
	}

	// 获取图片
	dir := fmt.Sprintf("%v/%v/%v",
		config.GetDeStr("avatar.dir", "./avatar"),
		id, "avatar.jpg")

	file, e := ioutil.ReadFile(dir)
	if e != nil {
		c.JSON(http.StatusOK, resp.NotExist.SetMessage(e))
		return

	}
	src := base64.StdEncoding.EncodeToString(file)

	c.JSON(http.StatusOK, resp.NewSucResp(src))

}
