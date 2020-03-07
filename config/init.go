package config

import (
	"log"
	"sync"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	defaultConfigFile = "config.json"
	defaultConfigPath = "./config/"
)

var onceInit sync.Once

// todo: attention!需要依赖配置的地方都必须手动调用一次该方法！
func Init() {
	onceInit.Do(func() {
		viper.SetConfigName(defaultConfigFile)
		viper.AddConfigPath(defaultConfigPath)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("[Load Config]failed:%v", err)
		}

		// 支持配置热更新
		viper.WatchConfig()
		viper.OnConfigChange(func(in fsnotify.Event) {
			log.Println("Config has been changed")

			if err := viper.ReadInConfig(); err != nil {
				log.Fatalf("[ReLoad Config]failed:%v", err)
			}

			log.Println("Config change finish ")
		})
	})
}
