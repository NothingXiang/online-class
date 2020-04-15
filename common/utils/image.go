package utils

import (
	"bytes"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"golang.org/x/image/bmp"
)

func EncodeToString(imagePath string) string {
	old_file, _ := os.Open(imagePath)

	result := bytes.NewBuffer(make([]byte, 0))
	// 识别图片类型
	file, image_type, _ := image.Decode(old_file)

	log.Printf("the image %v type is:%v", imagePath, image_type)

	// 获取图片的类型
	switch image_type {
	case `jpeg`:
		jpeg.Encode(result, file, nil)
	case `png`:
		png.Encode(result, file)
	case `gif`:
		gif.Encode(result, file, nil)
	case `bmp`:
		bmp.Encode(result, file)
	default:
		png.Encode(result, file)
	}
	old_file.Close()

	return result.String()
}
