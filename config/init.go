package config

import (
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/NothingXiang/online-class/common/utils"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	defaultConfigFile = "config.json"
	defaultConfigPath = "./config"
)

var (
	onceInit sync.Once

	GinMode string = "debug"
)

//  attention!需要依赖配置的地方最好手动调用一次该方法！
func Init() {
	onceInit.Do(func() {
		viper.SetConfigName(defaultConfigFile)
		viper.AddConfigPath(defaultConfigPath)
		if err := viper.ReadInConfig(); err != nil {
			log.Errorf("[Load Config]failed:%v", err)
		}

		// 支持配置热更新
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			log.Println("Config has been changed")
			if err := viper.ReadInConfig(); err != nil {
				log.Errorf("[ReLoad Config]failed:%v", err)
			}
			log.Println("Config change finish ")
		})

	})
}

// 从配置中获取值，如果获取不到则回传defaultValue
func GetDeStr(key, defVal string) string {
	var value string
	if value = viper.GetString(key); utils.IsEmptyString(value) {
		return defVal
	}
	return value
}
