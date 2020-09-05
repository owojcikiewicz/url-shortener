package app

import (
	"gorm.io/gorm"
	"log"
	"url-shortener/config"
)

type App struct {
	DB *gorm.DB
}

func (app *App) Run(config *config.Config) {
	err := app.InitializeDatabase(config)
	if err != nil {
		log.Fatal("A database error occurred: ", err)
	}

	err = app.InitializeRoutes(config.Port, config.Password)
	if err != nil {
		log.Fatal("A route error occurred: ", err)
	}
}
