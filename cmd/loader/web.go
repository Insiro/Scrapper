package loader

import (
    "Scrapper/internal/routes"
    "Scrapper/internal/util"
    "Scrapper/pkg/out"
    "github.com/gin-gonic/gin"
)

func Web(config *util.Config) {
    app := gin.Default()

    out.Table(config, "config")

    var basePath = config.BaseURL
    if basePath == "" {
        basePath = "/"
    } else if basePath[0] != '/' {
        basePath = "/" + basePath
    }
    route := app.Group(basePath)

    routes.ApiRoute(route)

    route.StaticFile("/", "./dist/index.html")
    route.Static("/assets", "./dist/assets")

    //	app.NoRoute(func(c *gin.Context) {
    //		c.File("./dist/index.html")
    //	})
    err := app.Run(":9000")
    if err != nil {
        panic(err)
    }
}
