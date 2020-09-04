package main

import (
	"url-shortener/app"
	"url-shortener/config"
)

func main() {
	cfg := config.LoadConfig()
	a := &app.App{}
	a.Run(cfg)
}