package notice

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(e *gin.Engine) {
	notice := e.Group("/notice")
	{
		notice.GET("/get", GetNoticeByClassPageable)

		notice.POST("/create", CreateNotice)

		notice.POST("/update", UpdateNotice)

		notice.GET("/get/template", GetTemplate)

	}
}
