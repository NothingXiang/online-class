package static

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/config"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
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
		static.POST("/upload/single", UploadSingleFile)

		static.POST("/upload/multi", UploadMutiFile)
	}

	static.Static("/get", config.GetDeStr("dir.base", BasePath))

}

// 上传单个文件
func UploadSingleFile(c *gin.Context) {
	// 1. get file data
	file, err := c.FormFile("file")
	if err != nil {
		resp.ErrJson(c, resp.ParamEmptyErr.NewErr(err))
		return
	}

	// 2. get file message
	fileType, suc := req.TryGetParam("type", c)
	if !suc || !CheckFileType(fileType) {
		resp.ErrJson(c, resp.ParamEmptyErr.NewErrStr("type param"))
		return
	}
	id := uuid.NewV4().String()

	//3. get dir ,make sure dir exist
	dir := fmt.Sprintf("%v/%v/%v",
		config.GetDeStr("dir.base", "./static"),
		fileType,
		id)
	os.MkdirAll(dir, os.ModeDir)

	// 4. save file

	dst := fmt.Sprintf("%v/%v", dir, filepath.Base(file.Filename))
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		resp.ErrJson(c, resp.UploadFileError.NewErr(err))
		return
	}

	resp.SucJson(c, dst)

}

func UploadMutiFile(c *gin.Context) {
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

	/* 4. save all file by goroutines*/
	paths := make([]string, len(files))

	// make sure goroutines finish
	var wg sync.WaitGroup
	wg.Add(len(files))

	// save file
	for index, file := range files {

		paths[index] = fmt.Sprintf("%v/%v", baseDir, filepath.Base(file.Filename))

		//
		go func(dst string) {
			err := c.SaveUploadedFile(file, dst)
			if err != nil {
				logrus.Error(err)
			}
			wg.Done()
		}(paths[index])
	}

	wg.Wait()

	// response file paths
	resp.SucJson(c, paths)
}
