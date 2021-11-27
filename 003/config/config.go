package config

import (
	"os"
)
import "github.com/spf13/viper"

func LoadConfig() error {
	env := os.Getenv("CONFIGFROM")
	if env == "local"{
		viper.SetConfigType("json")
		viper.SetConfigName("default")
		viper.AddConfigPath("/data/config")
	}
	viper.WatchConfig()
	err := viper.ReadInConfig()
	if err != nil{
		return err
	}
	return nil
}
