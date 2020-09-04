package main

import (
	"fmt"
	"url-shortener/config"
)

func main() {
	config := config.LoadConfig()
	fmt.Println(config)
}