package app

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"url-shortener/config"
)

type Link struct {
	gorm.Model
	ID string
	URL string
}

func (app *App) InitializeDatabase(config *config.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed: ", err)
		return err
	}

	err = db.AutoMigrate(&Link{})
	if err != nil {
		log.Fatal(err)
	}
	app.DB = db

	return nil
}