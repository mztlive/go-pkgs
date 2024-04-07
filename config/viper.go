package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// 初始化配置， 如果是生产环境则会覆盖配置文件里面的一些内容
func Initialize(configPath string) {
	// Set config file name based on environment
	configName := "local"

	// Read environment variable
	env := os.Getenv("APP_ENV")

	if env != "" {
		configName = env
	}

	// Set up viper
	viper.SetConfigName(configName)
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Fatal error reading config file: %s \n", err.Error())
	}

	// Override config values from environment if in prod
	if env != "local" {
		viper.Set("mongodb.uri", os.Getenv("MONGODB_URI"))
		viper.Set("app.host", os.Getenv("APP_HOST"))
		viper.Set("database.uri", os.Getenv("DATABASE_URI"))
	}
}
