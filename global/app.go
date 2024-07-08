package global

import (
	"gin_ready/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
}

var App = new(Application)
