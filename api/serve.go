/*
 *  提供api接口
 */
package api

import (
	"log"
	"net/http"

	"github.com/NothingXiang/online-class/api/class"
	"github.com/NothingXiang/online-class/api/courseware"
	"github.com/NothingXiang/online-class/api/homework"
	"github.com/NothingXiang/online-class/api/notice"
	"github.com/NothingXiang/online-class/api/static"
	"github.com/NothingXiang/online-class/api/survey"
	"github.com/NothingXiang/online-class/api/user"
	"github.com/NothingXiang/online-class/common/pkg"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

// 网络服务
func Serve(info pkg.Info) {

	// 设置路由运行模式（debug or release）
	gin.SetMode(viper.GetString("gin.mode"))

	engine := gin.Default()

	// 注册中间件/路由
	engine.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, info)
	})

	static.ServeStatic(engine)

	user.RegisterRoutes(engine)
	class.RegisterRouters(engine)
	notice.RegisterRouters(engine)
	survey.RegisterRoutes(engine)
	homework.RegisterRouters(engine)
	courseware.RegisterRouters(engine)

	// 课件

	port := viper.GetString("http.port")
	if err := engine.Run(":" + port); err != nil {
		log.Printf("GIN Server Fail:%v", err)
	}
}
