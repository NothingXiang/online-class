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

		// 通过微信id检查账户是否已经存在
		user.POST("/get/account/wechat", GetAccountByWeChat)

		// 通过wechat code 登录，注意，会检查该微信账号是否存在，不存在则返回新创建的账号
		user.POST("/login/wechat", LoginByWeChat)

		//	注销用户
		//user.DELETE(":userid", Logout)

		// 上传头像
		user.POST("/avatar", UploadAvatar)

		// todo:gin server static
		//	获取头像
		user.GET("/avatar", GetAvatar)
	}
}

