package loader

import (
    "Scrapper/internal/app"
    "Scrapper/internal/util"
    "Scrapper/pkg/out"
    "github.com/gin-gonic/gin"
    "strings"
)

func Web(config *util.Config) {
    g := gin.Default()

    out.Table(config, "config")

    var basePath = config.BaseURL
    if basePath == "" {
        basePath = "/"
    } else if basePath[0] != '/' {
        basePath = "/" + basePath
    }
    route := g.Group(basePath)

    app.ApiRoute(route)

    route.StaticFile("", "./dist/index.html")
    route.StaticFile("/", "./dist/index.html")
    route.Static("/assets", "./dist/assets")

    // NoRoute 핸들러를 사용해 정의되지 않은 모든 경로에 대해 index.html을 반환
    g.NoRoute(func(c *gin.Context) {
        subRoute := strings.TrimPrefix(c.Request.URL.Path, basePath)
        if subRoute == "/api" || subRoute == "/static" {
            c.String(404, "Not Found")
            return
        }
        c.File("./dist/index.html")
    })

    err := g.Run(":9000")
    if err != nil {
        panic(err)
    }
}
