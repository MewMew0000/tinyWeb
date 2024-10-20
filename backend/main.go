package main

import (
	"backend/config"
	"backend/global"
	"backend/router"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Printf("config info: %v\n", config.AppConfig)
	config.InitDB()
	fmt.Printf("database info: %v\n", global.Db)
	config.InitRedis()
	fmt.Printf("redis info: %v\n", global.RedisDB)
	r := router.SetupRouter()
	port := config.AppConfig.App.Port
	r.Run(port)
}
