package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	user := e.Group("user")
	{
		// 注册
		user.POST("/create", CreateUser)
		// 登录
		user.POST("/login", Login)

		//	注销用户
		//user.DELETE(":userid", Logout)
	}
}
