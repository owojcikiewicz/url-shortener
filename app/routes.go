package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func (app *App) InitializeRoutes(port string) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/u/:id", func(c *gin.Context) {
		var link Link
		id := c.Param("id")

		err := app.DB.First(&link, id).Error
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(link)
	})

	router.Run(fmt.Sprintf(":%s", port))

	return nil
}