package courseware

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(e *gin.Engine) {
	courseware := e.Group("/courseware")
	{
		courseware.POST("/create", CreateCourseware)

		courseware.GET("/get/info", GetCoursewareInfo)

		courseware.GET("/list/info", ListCoursewareInfo)

		courseware.DELETE("/del", RemoveCourseware)
	}
}
