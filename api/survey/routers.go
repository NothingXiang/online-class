package survey

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(e *gin.Engine) {
	survey := e.Group("/survey")
	{
		// 创建问卷
		survey.POST("/create", CreateSurvey)

		// 获取问卷
		survey.GET("/get", GetSurvey)

		// 分页获取班级问卷
		survey.GET("/list", ListSurvey)

		// 删除某分问卷
		survey.DELETE("/delete", DeleteSurvey)

		answer := survey.Group("/answer")
		{
			// 创建答卷
			answer.POST("/create", CreateAnswer)

			// 获取统计数据
			answer.GET("/statistics", GetStatistics)
		}

	}
}
