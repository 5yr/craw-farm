package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	Setup()
}

var cfg *viper.Viper

func Setup() {
	var (
		err error
	)

	viper.SetConfigName("conf")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("toml")
	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf(`fatal error read config file: %s \n`, err))
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println(`config file changed`)
	})
}

func GetSiteConfig(site string, rawValue interface{}) (err error) {
	return viper.Sub("websites").UnmarshalKey(site, rawValue)
}
