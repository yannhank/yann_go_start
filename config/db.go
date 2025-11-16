package config

import (
	"fmt"
	"log"
	"sc_police_api/global"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		AppConfig.Database.Host,
		AppConfig.Database.User,
		AppConfig.Database.Password,
		AppConfig.Database.Dbname,
		AppConfig.Database.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v", err)
		fmt.Println("Error connecting to PostgreSQL")
	} else {
		fmt.Println("success connecting to PostgreSQL")
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
		fmt.Println("error getting database instance")
	} else {
		fmt.Println("success getting database instance")
	}

	// 连接池配置
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetConnMaxLifetime(time.Hour)

	global.Db = db
	return db
}
