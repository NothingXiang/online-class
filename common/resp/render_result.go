package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 返回一个json结果
func Json(c *gin.Context, result interface{}) {
	c.JSON(http.StatusOK, result)
}

func ErrJson(c *gin.Context, err error) {
	c.JSON(http.StatusOK, ErrResp(err))
}

func SucJson(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewSucResp(data))
}
