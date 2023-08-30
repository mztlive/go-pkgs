package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// 初始化配置的路径
func Initialize(configPath string) {

	// 根据环境变量设置配置文件名字
	// 例如：export ENV=dev
	// 配置文件名字为：dev.toml

	// 读取环境变量
	env := os.Getenv("APP_ENV")
	switch env {
	case "dev":
		viper.SetConfigName("dev")
	case "test":
		viper.SetConfigName("test")
	case "prod":
		viper.SetConfigName("prod")
	default:
		viper.SetConfigName("local")
	}

	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig() // 查找并读取配置文件
	if err != nil {             // 处理读取配置文件的错误
		log.Fatalf("Fatal error config file: %s \n", err.Error())
	}
}
