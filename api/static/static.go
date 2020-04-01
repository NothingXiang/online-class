package static

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/NothingXiang/online-class/common/req"
	"github.com/NothingXiang/online-class/common/resp"
	"github.com/NothingXiang/online-class/config"
	"github.com/gin-gonic/gin"
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

	// 2. get file message,include file type and file id
	fileType, suc := req.TryGetParam("type", c)
	if !suc || !CheckFileType(fileType) {
		resp.ErrJson(c, resp.ParamEmptyErr.NewErrStr("type param"))
		return
	}
	id, ok := req.TryGetParam("id", c)
	if !ok {
		resp.ErrJson(c, resp.ParamEmptyErr.NewErrStr("id param "))
		return
	}

	//3. get dir ,make sure dir exist
	dir := fmt.Sprintf("%v/%v/%v",
		config.GetDeStr("dir.base", "./static"),
		fileType,
		id)
	os.MkdirAll(dir, os.ModeDir)

	// 4. save file
	err = c.SaveUploadedFile(file,
		fmt.Sprintf("%v/%v", dir, filepath.Base(file.Filename)))
	if err != nil {
		resp.ErrJson(c, resp.UploadFileError.NewErr(err))
		return
	}

	resp.SucJson(c, nil)

}

func UploadMutiFile(c *gin.Context) {
	// 1. get file data
	form, _ := c.MultipartForm()

	files := form.File["files"]

	// 2. get file message,include file type and file id
	id, ok := req.TryGetParam("id", c)
	if !ok {
		resp.ErrJson(c, resp.ParamEmptyErr.NewErrStr("id param "))
		return
	}
	fileType, suc := req.TryGetParam("type", c)
	if !suc || !CheckFileType(fileType) {
		resp.ErrJson(c, resp.ParamEmptyErr.NewErrStr("type param"))
		return
	}

	//3. get dir ,make sure dir exist
	dir := fmt.Sprintf("%v/%v/%v",
		config.GetDeStr("dir.base", "./static"),
		fileType,
		id)
	os.MkdirAll(dir, os.ModeDir)

	// 4. save all file
	for _, file := range files {
		err := c.SaveUploadedFile(file,
			fmt.Sprintf("%v/%v", dir, filepath.Base(file.Filename)))
		if err != nil {
			resp.ErrJson(c, resp.UploadFileError.NewErr(err))
			return
		}

	}
	resp.SucJson(c, nil)
}
