package main

import (
	"Scrapper/internal/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	app := gin.Default()
	godotenv.Load()

	basePath := os.Getenv("SCRAPER_BASE_PATH")
	if basePath == "" {
		basePath = "/"
	} else if basePath[0] != '/' {
		basePath = "/" + basePath
	}
	fmt.Print(basePath)
	route := app.Group(basePath)

	routes.ApiRoute(route)

	route.StaticFile("/", "./dist/index.html")
	route.Static("/assets", "./dist/assets")

	//	app.NoRoute(func(c *gin.Context) {
	//		c.File("./dist/index.html")
	//	})
	app.Run(":9000")
	
}
