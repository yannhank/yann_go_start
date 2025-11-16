package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	App struct {
		Name string
		Port string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		Dbname   string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")    // 配置文件名
	viper.SetConfigType("yaml")      // 配置文件类型
	viper.AddConfigPath("./config/") // 配置文件路径

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %v", err)
		fmt.Println("reading config error")
	} else {
		fmt.Println("reading config success")
	}

	AppConfig = &Config{}

	err := viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
		fmt.Println("Unable to decode error")
	} else {
		fmt.Println("Unable to decode success", AppConfig)
	}

	InitDb()
}
