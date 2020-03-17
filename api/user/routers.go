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
		user.POST("/login", LoginByPhonePwd)

		//	注销用户
		//user.DELETE(":userid", Logout)

		// 上传头像
		user.POST("/avatar", UploadAvatar)

		// todo:gin server static
		//	获取头像
		user.GET("/avatar", GetAvatar)
	}
}
