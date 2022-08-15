package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var Config ProjectConfig

type ProjectConfig struct {
	Port      int
	UrlPrefix string
	Log       LogConfig
	Time      TimeConfig
}

type LogConfig struct {
	WriteFile bool
	FileDir   string
	FileName  string
}

type TimeConfig struct {
	TimeZoneStr string
	TimeZone    *time.Location
}

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs/")

	// Load config file
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error read config file: %s \n", err))
	}
	err = viper.Unmarshal(&Config)
	if err != nil {
		panic(err)
	}

	if Config.Time.TimeZoneStr == "" {
		Config.Time.TimeZoneStr = "Asia/Shanghai"
	}
	Config.Time.TimeZone, err = time.LoadLocation(Config.Time.TimeZoneStr)
	if err != nil {
		panic(err)
	}

}
