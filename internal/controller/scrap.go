package controller

import (
    "Scrapper/internal/dto"
    "Scrapper/internal/entity"
    "Scrapper/internal/service"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type Scrap struct {
    scrap service.Scrap
    route *gin.RouterGroup
}

var _ IController = (*Scrap)(nil)

func ScrapController(scrap service.Scrap, parent *gin.RouterGroup) Scrap {
    cont := Scrap{scrap: scrap}
    cont.Init(parent)
    return cont
}

func (r *Scrap) ParseUrl(c *gin.Context) {
    var input = dto.URLInput{}
    if err := c.ShouldBindJSON(&input); err != nil {
        _ = c.Error(err)
        return
    }
    scrap, err := r.scrap.Scrap(input.Url)
    if scrap == nil {
        _ = c.Error(err)
        return
    }

    if err != nil {
        c.JSON(http.StatusOK, dto.NewScrap(scrap, nil))
        return
    }

    c.JSON(http.StatusCreated, dto.NewScrap(scrap, []entity.Tag{}))
}

func (r *Scrap) ReScrap(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        _ = c.Error(err)
        return
    }
    scrap, err := r.scrap.ReScrap(id)
    if err != nil {
        _ = c.Error(err)
        return
    }

    c.JSON(http.StatusAccepted, dto.NewScrap(scrap, nil))
    return
}

func (r *Scrap) Detail(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    scrap, err := r.scrap.Get(id)
    if err != nil {
        _ = c.Error(err)
        return
    }

    c.JSON(http.StatusOK, dto.NewScrap(scrap, scrap.Tags))
}

func (r *Scrap) List(c *gin.Context) {
    input := dto.ListScrap{}

    err := c.BindQuery(&input)
    count, err := r.scrap.Count()
    if err != nil {
        _ = c.Error(err)
        return
    }

    scraps, err := r.scrap.List(input.Offset, input.Limit, input.Pined)
    if err != nil {
        _ = c.Error(err)
        return
    }
    dtos := make([]dto.Scrap, len(scraps))
    for i, scrap := range scraps {
        dtos[i] = dto.NewScrap(&scrap, scrap.Tags)
    }

    c.JSON(http.StatusOK, gin.H{"list": dtos, "count": count})
    return
}

func (r *Scrap) Update(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        _ = c.Error(err)
        return
    }
    var data = dto.ScrapUpdate{}
    err = c.ShouldBindBodyWithJSON(&data)
    if err != nil {
        _ = c.Error(err)
        return
    }
    scrap, err := r.scrap.Update(id, data)
    if err != nil {
        _ = c.Error(err)
        return
    }
    c.JSON(http.StatusOK, dto.NewScrap(scrap, scrap.Tags))
    return

}

func (r *Scrap) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        _ = c.Error(err)
        return
    }
    if err = r.scrap.Delete(id); err != nil {
        _ = c.Error(err)
        return
    }
    c.String(http.StatusOK, http.StatusText(http.StatusOK))
}

func (r *Scrap) Init(parent *gin.RouterGroup) IController {
    var g = parent.Group("/scraps")
    r.route = g
    g.GET("", r.List)
    g.POST("", r.ParseUrl)

    g.GET(":id", r.Detail)
    g.POST(":id", r.ReScrap)
    g.PATCH(":id", r.Update)

    return r
}
