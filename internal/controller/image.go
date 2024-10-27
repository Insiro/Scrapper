package controller

import (
    "Scrapper/internal/model/dto"
    "Scrapper/internal/repository"
    "github.com/gin-gonic/gin"
    "net/http"
)

type ImageController struct {
    repo repository.ImageRepository
}

func (i *ImageController) Delete(c *gin.Context, id uint) *gin.Error {
    if err := i.repo.Delete([]uint{id}); err != nil {
        return c.Error(err)
    }
    c.String(http.StatusOK, http.StatusText(http.StatusOK))

    return nil
}

func (i *ImageController) DeleteList(c *gin.Context, list dto.ImageDelete) *gin.Error {
    if err := i.repo.Delete(list.Images); err != nil {
        return c.Error(err)
    }
    c.String(http.StatusOK, http.StatusText(http.StatusOK))

    return nil
}
