package controller

import (
    "Scrapper/internal/dto"
    "Scrapper/internal/entity"
    repos "Scrapper/internal/repository"
    "Scrapper/internal/scrapper"
    "Scrapper/internal/util"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type Scrap struct {
    repo    repos.Scrap
    imgRepo repos.Image
    route   *gin.RouterGroup
    util.Config
}

var _ IController = (*Scrap)(nil)

func ScrapController(repo repos.Scrap, imgRepo repos.Image, parent *gin.RouterGroup, config util.Config) Scrap {
    cont := Scrap{repo, imgRepo, nil, config}
    cont.Init(parent)
    return cont
}

func (r *Scrap) ParseUrl(c *gin.Context) {
    var input = dto.URLInput{}
    if err := c.ShouldBindJSON(&input); err != nil {
        _ = c.Error(err)
        return
    }

    scraper, err := scrapper.Factory(input.Url, nil, nil)
    if err != nil {
        _ = c.Error(err)
        return
    }
    args := scraper.Args
    exist, err := r.repo.GetBySourceId(scraper.PageType, args.Key)
    if err == nil {
        data := dto.NewScrap(exist, []entity.Tag{})
        c.JSON(http.StatusOK, data)
        return
    }
    scrapCreate, err := scraper.Scrap(nil, r.Media)
    if err != nil {
        _ = c.Error(err)
        return
    }
    saved, err := r.repo.Create(scrapCreate)
    if err != nil {
        _ = c.Error(err)
        return
    }
    err = r.imgRepo.SaveImage(saved.ID, scrapCreate.ImageNames)
    if err != nil {
        _ = c.Error(err)
        return
    }

    c.JSON(http.StatusCreated, dto.NewScrap(saved, []entity.Tag{}))
    return
}

func (r *Scrap) ReScrap(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        _ = c.Error(err)
        return
    }

    found, err := r.repo.GetScrap(id)
    if err != nil {
        _ = c.Error(err)
        return
    }
    scr, _ := scrapper.Factory(found.Url(), &found.Source, nil)
    data, err := scr.Scrap(nil, r.Media)
    if err != nil {
        _ = c.Error(err)
        return
    }
    updateData := dto.Create2Update(data)

    result, err := r.repo.Update(found, updateData)
    if err != nil {
        _ = c.Error(err)
        return
    }
    err = r.imgRepo.DeleteByScrapId(id)
    err2 := r.imgRepo.SaveImage(id, data.ImageNames)
    if err != nil {
        _ = c.Error(err)
        return
    }
    if err2 != nil {
        _ = c.Error(err)
        return
    }

    c.JSON(http.StatusAccepted, result)
    return
}

func (r *Scrap) Detail(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        _ = c.Error(err)
        return
    }

    scrap, err := r.repo.GetScrap(id)
    if err != nil {
        _ = c.Error(err)
        return
    }
    //TODO: load Image
    tags, _ := r.repo.GetTags(id, "")
    c.JSON(http.StatusOK, dto.NewScrap(scrap, tags))

    return
}

func (r *Scrap) List(c *gin.Context) {
    input := dto.ListScrap{}

    err := c.BindQuery(&input)

    scraps, err := r.repo.ListScrap(input.Offset, input.Limit, input.Pined)
    if err != nil {
        _ = c.Error(err)
        return
    }
    count, err := r.repo.CountScrap()
    if err != nil {
        _ = c.Error(err)
        return
    }
    dtos := make([]dto.Scrap, len(scraps))
    for i, scrap := range scraps {
        tag, _ := r.repo.GetTags(scrap.ID, "")
        dtos[i] = dto.NewScrap(scrap, tag)
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

    scrap, err := r.repo.GetScrap(id)
    if err != nil {
        _ = c.Error(err)
        return
    }
    scrap, err = r.repo.Update(scrap, data)
    if err != nil {
        _ = c.Error(err)
        return
    }
    tags, _ := r.repo.GetTags(id, "")

    c.JSON(http.StatusOK, dto.NewScrap(scrap, tags))
    return

}

func (r *Scrap) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        _ = c.Error(err)
        return
    }

    if err := r.repo.DeleteScrap(id); err != nil {
        _ = c.Error(err)
        return
    }

    if err = r.imgRepo.DeleteByScrapId(id); err != nil {
        _ = c.Error(err)
        return
    }

    c.String(http.StatusOK, http.StatusText(http.StatusOK))
    return
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
