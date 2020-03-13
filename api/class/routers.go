package class

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouters(e *gin.Engine) {
	class := e.Group("class")
	{
		class.POST("/create",CreateClass)

		class.GET("get",GetClass)
	}
}
