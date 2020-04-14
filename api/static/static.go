package static

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/config"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	// 文件类型
	Avatar = "avatar"

	Homework = "homework"

	Courseware = "courseware"

	//	文件存储基本目录
	BasePath = "./static"
)

func CheckFileType(t string) bool {
	return t == Avatar || t == Homework || t == Courseware
}

func ServeStatic(e *gin.Engine) {

	static := e.Group("/static")
	{
		// 上传文件
		static.POST("/upload", UploadFiles)
	}

	e.Static(viper.GetString("file.path"), config.GetDeStr("dir.base", BasePath))

}

// 上传文件
func UploadFiles(c *gin.Context) {
	/*1. get file data*/
	form, _ := c.MultipartForm()

	files := form.File["files"]

	/*2. get file type*/
	fileType, suc := req.TryGetParam("type", c)
	if !suc || !CheckFileType(fileType) {
		resp.ErrJson(c, resp.ParamEmptyErr.NewErrStr("type param"))
		return
	}

	/*3. generate baseDir ,make sure baseDir exist*/
	id := uuid.NewV4().String()
	baseDir := fmt.Sprintf("%v/%v/%v",
		config.GetDeStr("baseDir.base", "./static"),
		fileType,
		id)
	os.MkdirAll(baseDir, os.ModeDir)

	/* 4. save all file by goroutines ,and return the path finally*/

	accessPaths := make([]string, len(files))

	// make sure goroutines finish
	var wg sync.WaitGroup
	wg.Add(len(files))

	// save file
	for index, file := range files {

		// local path
		localPath := fmt.Sprintf("%v/%v", baseDir, filepath.Base(file.Filename))

		// outside access path
		accessPaths[index] = strings.ReplaceAll(localPath,
			config.GetDeStr("baseDir.base", "./static"),
			fmt.Sprintf("%v%v", viper.GetString("file.addr"), viper.GetString("file.path")),
		)

		// save file by goroutines
		go func(dst string) {
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				logrus.Error(err)
			}
			wg.Done()
		}(localPath)

	}

	wg.Wait()

	// response file accessPaths
	resp.SucJson(c, accessPaths)
}
