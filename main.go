/*
程序入口main
*/
package main

import (
	"log"
	"time"

	"github.com/NothingXiang/online-class/api"
	"github.com/NothingXiang/online-class/common/pkg"
	"github.com/NothingXiang/online-class/config"
)

var PkgInfo = pkg.Info{
	AppName:   "online-class",
	Version:   "0.1.0",
	StartTime: time.Now(),
}

func main() {

	// 1. load common line args
	// todo:may be can update to cobra
	//flag.Parse()

	//2. recover
	defer func() {
		if r := recover(); r != nil {
			log.Println(PkgInfo.AppName, "Process Stop : ", r)
		} else {
			log.Println(PkgInfo.AppName, "Process Stop")
		}
	}()

	// 3. load config
	config.Init()

	// 4. set routers
	api.Serve(PkgInfo)

}
