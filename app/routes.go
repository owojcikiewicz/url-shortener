package app

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/ventu-io/go-shortid"
	"log"
	"net/url"
)

type Data struct {
	URL  string `json:"url"`
	Slug string `json:"slug"`
}

func (app *App) GenerateToken() (token string, error error) {
	sid, err := shortid.Generate()

	return sid, err
}

func (app *App) IsValidURL(urls string) (valid bool) {
	_, err := url.ParseRequestURI(urls)
	if err != nil {
		return false
	}

	u, err := url.Parse(urls)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

func (app *App) InitializeRoutes(port string, password string) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization"},
	}))
	router.Use(static.Serve("/", static.LocalFile("ui/", false)))

	router.POST("/create", func(c *gin.Context) {
		key := c.Request.Header.Get("Authorization")
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

		validURL := app.IsValidURL(data.URL)
		if validURL == false {
			c.Status(400)
			return
		}

		err = app.DB.Where("id = ?", data.Slug).First(&link).Error
		if err == nil {
			c.Status(409)
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

	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}