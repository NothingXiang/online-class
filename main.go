package main

import (
	"log"
	"time"

	"online-class/common/pkg"
	"online-class/config"
)

var PkgInfo = pkg.Info{
	AppName:   "online-class",
	Version:   "0.1.0",
	StartTime: time.Now(),
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(PkgInfo.AppName, "Process Stop : ", r)
		} else {
			log.Println(PkgInfo.AppName, "Process Stop")
		}
	}()

	config.Init()
}
