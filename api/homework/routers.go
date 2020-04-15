package homework

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(e *gin.Engine) {
	homework := e.Group("homework")
	{
		homework.GET("/get", GetHomework)

		homework.GET("/list", ListHomework)

		homework.POST("/create", CreateHomework)

		homework.POST("/update", UpdateHomework)

		homework.DELETE("/delete", RemoveHomework)

		// 每项作业的已读列表
		read := homework.Group("/read")
		{
			read.GET("/list", GetReadList)

			read.POST("/add", AddReadList)
		}
	}
}

