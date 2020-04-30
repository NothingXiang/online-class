/*
 用户可以有多种登录/注册方式：手机号，邮箱验证等，但后续行为都只校验id和password
*/

package user

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	user := e.Group("user")
	{
		// 注册
		user.POST("/create", CreateUserByPhonePwd)
		// 登录
		user.POST("/login/phone", LoginByPhonePwd)

		// 通过微信id登录
		user.GET("/login/wechat", LoginByWeChatCode)

		// 通过WeChat 创建账号，注意，检查该微信账号是否存在(openID)
		user.POST("/create/wechat", CreateByWeChat)

		user.POST("/update", UpdateUser)

		user.GET("/get",GetUser)

		//	注销用户
		//user.DELETE(":userid", Logout)

	}
}
