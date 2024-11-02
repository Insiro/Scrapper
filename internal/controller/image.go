package controller

import (
    "Scrapper/internal/app"
    "Scrapper/internal/dto"
    "Scrapper/internal/repository"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type Image struct {
    repo   repository.Image
    route  *gin.RouterGroup
    config app.Config
}

func ImageController(repo repository.Image, parent *gin.RouterGroup, config app.Config) Image {
    cont := Image{repo, nil, config}
    cont.Init(parent)
    return cont
}

func (i *Image) Init(parent *gin.RouterGroup) IController {

    var g = parent.Group("/images")
    i.route = g

    g.DELETE("", i.DeleteList)

    g.DELETE(":id", i.Delete)

    return i
}

var _ IController = (*Image)(nil)

func (i *Image) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Request.PathValue("id"))
    if err != nil {
        _ = c.Error(err)
        return
    }
    if err := i.repo.Delete([]int{id}); err != nil {
        _ = c.Error(err)
        return
    }
    c.String(http.StatusOK, http.StatusText(http.StatusOK))

    return
}

func (i *Image) DeleteList(c *gin.Context) {
    var input = dto.ImageDelete{}
    if err := c.ShouldBindJSON(&input); err != nil {
        _ = c.Error(err)
        return
    }
    if err := i.repo.Delete(input.Images); err != nil {
        _ = c.Error(err)
        return
    }
    c.String(http.StatusOK, http.StatusText(http.StatusOK))

    return
}
