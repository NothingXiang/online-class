package user

import (
	"net/http"

	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/models"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

}

func CreateUser(c *gin.Context) {

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusOK, resp.ErrResp(resp.ParamFmtErr.SetMsg(err)))
		return
	}

	//	 todo:入库操作...


	c.JSON(http.StatusOK, resp.Suc.SetData(user))
}
