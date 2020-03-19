package req

import (
	"net/http"
	"strconv"

	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/common/utils"
	"github.com/gin-gonic/gin"
)

// 检查入参中是否有空白字段，有则同时返回json错误
func CheckEmpty(c *gin.Context, params ...string) bool {
	for _, p := range params {
		if utils.IsEmptyString(p) {
			c.JSON(http.StatusOK, resp.ParamEmptyErr)
			return false
		}

	}
	return true
}

// 尝试从请求中获取参数
func TryGetParam(key string, c *gin.Context) (val string, suc bool) {

	// ?key1="..."&key2="...."
	if val = c.Query(key); val != "" {
		return val, true
	} else if val = c.Param(key); val != "" {
		return val, true
	} else if val = c.PostForm(key); val != "" {
		return val, true
	}
	return "", false
}

func TryGetInt(key string, c *gin.Context) (val int, suc bool) {
	var v string
	v, suc = TryGetParam(key, c)
	if !suc {
		return 0, false
	}

	var err error

	if val, err = strconv.Atoi(v); err != nil {
		return 0, false
	}

	return val, true
}
