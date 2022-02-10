package bootstrap

import (
	"fmt"
	"gin-demo/global"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	default_config_file = "config.yaml"
)

func InitConfig() *viper.Viper {
	config := default_config_file
	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		config = configEnv
	}

	// 初始化viper
	v := viper.New()

	v.SetConfigFile(config)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config failed: %s", err))
	}
	// 监听配置文件
	v.WatchConfig()

	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed: ", in.Name)
		if err := v.Unmarshal(&global.App.Config); err != nil {
			fmt.Errorf("reload config failed: %s", err)
		}
	})

	if err := v.Unmarshal(&global.App.Config); err != nil {
		fmt.Errorf("set config failed: %s", err)
	}
	return v
}
