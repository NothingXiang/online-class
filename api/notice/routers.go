package notice

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(e *gin.Engine) {
	notice := e.Group("/notice")
	{
		notice.GET("/list", GetNoticeByClassPageable)

		// 添加某条通知的已读列表
		notice.POST("/add/read", AddNoticeRead)

		notice.POST("/create", CreateNotice)

		notice.POST("/update", UpdateNotice)

		notice.GET("/get/template", GetTemplate)

	}
}
