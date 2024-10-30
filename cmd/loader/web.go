package loader

import (
    "Scrapper/internal/app"
    "Scrapper/internal/controller"
    "Scrapper/internal/repository"
    "Scrapper/internal/util"
    "Scrapper/pkg/out"
    "fmt"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "strings"
)

func Web(config *util.Config, db *gorm.DB) {
    g := gin.Default()

    out.Table(config, "config")

    var basePath = config.BaseURL
    if basePath != "" && basePath[0] == '/' {
        basePath = basePath[1:]
    }

    route := g.Group(basePath)
    apiRoute := app.ApiRoute(route)

    controller.NewScrapController(
        repository.NewScrapRepository(db),
        repository.NewImageRepository(db, "", "./ss/exp"),
        apiRoute,
    )

    route.StaticFile("", "./dist/index.html")
    route.StaticFile("/", "./dist/index.html")
    route.Static("/assets", "./dist/assets")

    // NoRoute 핸들러를 사용해 정의되지 않은 모든 경로에 대해 index.html을 반환
    g.NoRoute(func(c *gin.Context) {

        subRoute := strings.TrimPrefix(c.Request.URL.Path, "/"+basePath)
        fmt.Println(c.Request.URL.Path)
        fmt.Println(basePath)
        fmt.Println(subRoute)

        if strings.HasPrefix(subRoute, "/api") || strings.HasPrefix(subRoute, "/static") {
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
