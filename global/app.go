package global

import (
	"gin-demo/config"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
}

var App = new(Application)
