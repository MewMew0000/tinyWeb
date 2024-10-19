package main

import (
	"backend/config"
	"fmt"
)

func main() {
	config.InitConfig()
	fmt.Printf("config info: %v\n", config.AppConfig)
}
