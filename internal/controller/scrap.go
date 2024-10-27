package controller

import (
    "Scrapper/internal/model/dto"
    "Scrapper/internal/model/entity"
    "Scrapper/internal/repository"
    "Scrapper/internal/scrapper"
    "github.com/gin-gonic/gin"
    "net/http"
)

type ScrapController struct {
    repo    repository.ScrapRepository
    imgRepo repository.ImageRepository
}

func (r *ScrapController) ParseUrl(c *gin.Context, input dto.URLInput) *gin.Error {
    scraper, err := scrapper.Factory(input.Url, nil, nil)
    args := scraper.Args
    exist, err := r.repo.GetBySourceId(scraper.PageType, args.Key)
    if err != nil {
        data := dto.NewScrapResponse(exist, []entity.Tag{})
        c.JSON(http.StatusOK, data)
        return nil
    }
    scrapCreate, err := scraper.Scrap(nil)
    if err != nil {
        return c.Error(err)
    }
    saved, err := r.repo.Create(scrapCreate)
    err = r.imgRepo.SaveImage(saved.ID, scrapCreate.ImageNames)
    if err != nil {
        return c.Error(err)
    }

    c.JSON(http.StatusCreated, dto.NewScrapResponse(exist, []entity.Tag{}))
    return nil
}

func (r *ScrapController) ReScrap(c *gin.Context, id uint) *gin.Error {
    found, err := r.repo.GetScrap(id)
    if err != nil {
        return c.Error(err)
    }
    scr, _ := scrapper.Factory(found.Url(), &found.Source, nil)
    data, err := scr.Scrap(nil)
    if err != nil {
        return c.Error(err)
    }
    updateData := dto.Create2Update(data)

    result, err := r.repo.Update(found, updateData)
    if err != nil {
        return c.Error(err)
    }
    err = r.imgRepo.DeleteByScrapId(id)
    err2 := r.imgRepo.SaveImage(id, data.ImageNames)
    if err != nil {
        return c.Error(err)
    }
    if err2 != nil {
        return c.Error(err)
    }

    c.JSON(http.StatusAccepted, result)
    return nil
}

func (r *ScrapController) Detail(c *gin.Context, id uint) error {
    scrap, err := r.repo.GetScrap(id)
    if err != nil {
        return c.Error(err)
    }
    //TODO: load Images
    tags, _ := r.repo.GetTags(id, "")
    c.JSON(http.StatusOK, dto.NewScrapResponse(scrap, tags))

    return nil
}

func (r *ScrapController) List(c *gin.Context, offset int, limit int, pined bool) *gin.Error {
    scraps, err := r.repo.ListScrap(offset, limit, pined)
    count, err2 := r.repo.CountScrap()
    if err != nil {
        return c.Error(err)
    }
    if err2 != nil {
        return c.Error(err)
    }
    dtos := make([]dto.ScrapResponse, len(scraps))
    for i, scrap := range scraps {
        tag, _ := r.repo.GetTags(scrap.ID, "")
        dtos[i] = dto.NewScrapResponse(scrap, tag)
    }

    c.JSON(http.StatusOK, gin.H{"list": dtos, "count": count})
    return nil
}

func (r *ScrapController) Update(c *gin.Context, id uint, data dto.ScrapUpdate) error {
    scrap, err := r.repo.GetScrap(id)
    if err != nil {
        return c.Error(err)
    }
    scrap, err = r.repo.Update(scrap, data)
    if err != nil {
        return c.Error(err)
    }
    tags, _ := r.repo.GetTags(id, "")

    c.JSON(http.StatusOK, dto.NewScrapResponse(scrap, tags))
    return nil

}

func (r *ScrapController) Delete(c *gin.Context, id uint) error {
    err := r.repo.DeleteScrap(id)
    if err != nil {
        return c.Error(err)
    }
    err = r.imgRepo.DeleteByScrapId(id)
    if err != nil {
        return c.Error(err)
    }
    c.String(http.StatusOK, http.StatusText(http.StatusOK))
    return nil
}
