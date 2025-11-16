package main

import (
	"log"
	"sc_police_api/config"
	"sc_police_api/controllers"
	"sc_police_api/router"
)

func main() {
	config.InitConfig()

	r := router.SetupRouter()

	// 初始化数据库...
	if err := controllers.MigrateTables(); err != nil {
		log.Fatalf("数据库表迁移失败: %v", err)
	}

	// 映射 /uploads 到 ./uploads
	r.Static("/uploads", "./uploads")

	port := config.AppConfig.App.Port
	if port == "" {
		port = ":8088"
	}
	r.Run(port)
}
