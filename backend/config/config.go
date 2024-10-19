package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string
		Port string
	}
	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenConns int
	}
	Redis struct {
		Addr     string
		DB       int
		Password string
	}
}

var AppConfig *Config

func InitConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	AppConfig = &Config{}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("reading config failed: %v", err)
	}
	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatalf("unmarshal config failed: %v", err)
	}
}
