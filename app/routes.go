package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ventu-io/go-shortid"
)

type Data struct {
	URL  string `json:"url"`
	Slug string `json:"slug"`
}

func (app *App) GenerateToken() (token string, error error) {
	sid, err := shortid.Generate()

	return sid, err
}

func (app *App) InitializeRoutes(port string, password string) error {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.POST("/create", func(c *gin.Context) {
		key := c.Request.Header.Get("Authentication")
		data := Data{}
		link := Link{}

		if key != password {
			c.Status(401)
			return
		}

		err := c.BindJSON(&data)
		if err != nil {
			c.Status(400)
			return
		}

		err = app.DB.Where("id = ?", data.Slug).First(&link).Error
		if err == nil {
			c.String(400, "Slug In Use")
			return
		}

		if data.Slug == "" {
			id, err := app.GenerateToken()
			if err != nil {
				c.Status(500)
				return
			}

			data.Slug = id
		}

		link = Link{ID: data.Slug, URL: data.URL, Views: 0}
		err = app.DB.Create(&link).Error
		if err != nil {
			c.Status(500)
		}

		c.String(200, data.Slug)
	})

	router.GET("/:id", func(c *gin.Context) {
		link := Link{}
		id := c.Param("id")

		err := app.DB.Where("id = ?", id).First(&link).Error
		if err != nil {
			c.String(404, "Link Not Found")
			return
		}

		c.Redirect(307, link.URL)
		link.Views = link.Views + 1
		app.DB.Save(&link)
	})

	router.Run(fmt.Sprintf(":%s", port))
	return nil
}