package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (app *App) InitializeRoutes(port string) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/upload", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			return
		}
		fmt.Println(data)
	})

	router.GET("/:id", func(c *gin.Context) {
		link := Link{}
		id := c.Param("id")

		err := app.DB.Where("id = ?", id).First(&link).Error
		if err != nil {
			c.String(404, "404 - Link Not Found")
			return
		}

		c.Redirect(307, link.URL)
		link.Views = link.Views + 1
		app.DB.Save(&link)
	})

	router.Run(fmt.Sprintf(":%s", port))

	return nil
}