package config

import (
	"flag"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/NothingXiang/online-class/common/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	defaultFile = "config.json"
	defaultPath = "./config"
)

var (

	// 指定配置文件的路径
	Path = flag.String(
		"cp",
		defaultPath,
		"specify config file path")

	// 保证读取配置初始化只执行一次
	onceInit sync.Once
)

//  attention!  api 包之外需要依赖配置的地方，最好手动调用一次该方法！
func Init() {

	onceInit.Do(func() {

		viper.AddConfigPath(*Path)
		if err := viper.ReadInConfig(); err != nil {
			log.Errorf("[Load Config]failed:%v", err)
		}
		log.Printf("[Load Config] load %v success", *Path)

		// 支持配置热更新
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			log.Println("Config has been changed")

			if err := viper.ReadInConfig(); err != nil {
				log.Errorf("[ReLoad Config]failed:%v\n", err)
				return
			}

			log.Println("Config change finish ")
		})

	})
}

// 从配置中获取值，如果获取不到则回传defaultValue
func GetDeStr(key, defVal string) string {
	value := viper.GetString(key)
	if utils.IsEmptyString(value) {
		return defVal
	}
	return value
}
