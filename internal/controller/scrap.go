package controller

import (
    "Scrapper/internal/model/dto"
    "Scrapper/internal/model/entity"
    "Scrapper/internal/repository"
    "Scrapper/internal/scrapper"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type ScrapController struct {
    repo    repository.ScrapRepository
    imgRepo repository.ImageRepository
    parent  *gin.IRouter
}

func NewScrapController(repo repository.ScrapRepository, imgRepo repository.ImageRepository, parent gin.IRouter) ScrapController {
    controller := ScrapController{repo, imgRepo, &parent}
    return controller.Init()
}

func (r *ScrapController) ParseUrl(c *gin.Context) {
    var input = dto.URLInput{}
    if err := c.ShouldBindJSON(&input); err != nil {
        _ = c.Error(err)
        return
    }

    scraper, err := scrapper.Factory(input.Url, nil, nil)
    args := scraper.Args
    exist, err := r.repo.GetBySourceId(scraper.PageType, args.Key)
    if err != nil {
        data := dto.NewScrapResponse(exist, []entity.Tag{})
        c.JSON(http.StatusOK, data)
        return
    }
    scrapCreate, err := scraper.Scrap(nil)
    if err != nil {
        _ = c.Error(err)
        return
    }
    saved, err := r.repo.Create(scrapCreate)
    err = r.imgRepo.SaveImage(saved.ID, scrapCreate.ImageNames)
    if err != nil {
        _ = c.Error(err)
        return
    }

    c.JSON(http.StatusCreated, dto.NewScrapResponse(exist, []entity.Tag{}))
    return
}

func (r *ScrapController) ReScrap(c *gin.Context) {
    id, err := strconv.Atoi(c.Request.PathValue("id"))
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
    data, err := scr.Scrap(nil)
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

func (r *ScrapController) Detail(c *gin.Context) {
    id, err := strconv.Atoi(c.Request.PathValue("id"))
    if err != nil {
        _ = c.Error(err)
        return
    }

    scrap, err := r.repo.GetScrap(id)
    if err != nil {
        _ = c.Error(err)
        return
    }
    //TODO: load Images
    tags, _ := r.repo.GetTags(id, "")
    c.JSON(http.StatusOK, dto.NewScrapResponse(scrap, tags))

    return
}

func (r *ScrapController) List(c *gin.Context) {
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
    dtos := make([]dto.ScrapResponse, len(scraps))
    for i, scrap := range scraps {
        tag, _ := r.repo.GetTags(scrap.ID, "")
        dtos[i] = dto.NewScrapResponse(scrap, tag)
    }

    c.JSON(http.StatusOK, gin.H{"list": dtos, "count": count})
    return
}

func (r *ScrapController) Update(c *gin.Context) {
    id, err := strconv.Atoi(c.Request.PathValue("id"))
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

    c.JSON(http.StatusOK, dto.NewScrapResponse(scrap, tags))
    return

}

func (r *ScrapController) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Request.PathValue("id"))
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

func (r *ScrapController) Init() ScrapController {
    var g = (*r.parent).Group("/scraps")
    g.GET("", r.List)
    g.POST("", r.ParseUrl)
    return *r
}
