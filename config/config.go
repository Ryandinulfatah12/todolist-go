package config

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "todolist-go/models"
)

var DB *gorm.DB

func ConnectDB() {
    dsn := "root:@tcp(127.0.0.1:3306)/tododb?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrate the schema
    DB.AutoMigrate(&models.Todo{}, &models.Task{})
}
