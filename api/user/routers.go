package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	user := e.Group("user")
	{
		// 注册
		user.POST("/create", CreateUserByPwd)
		// 登录
		user.POST("/login", LoginByPwd)

		//	注销用户
		//user.DELETE(":userid", Logout)

		// 上传头像
		user.POST("/avatar", UploadAvatar)

		//	获取头像
		user.GET("/avatar", GetAvatar)
	}
}
