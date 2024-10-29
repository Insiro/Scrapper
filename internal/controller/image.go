package controller

import (
    "Scrapper/internal/model/dto"
    "Scrapper/internal/repository"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type ImageController struct {
    repo repository.ImageRepository
}

func (i *ImageController) Delete(c *gin.Context) *gin.Error {
    id, err := strconv.Atoi(c.Request.PathValue("id"))
    if err != nil {
        return c.Error(err)
    }
    if err := i.repo.Delete([]int{id}); err != nil {
        return c.Error(err)
    }
    c.String(http.StatusOK, http.StatusText(http.StatusOK))

    return nil
}

func (i *ImageController) DeleteList(c *gin.Context) *gin.Error {
    var input = dto.ImageDelete{}
    if err := c.ShouldBindJSON(&input); err != nil {
        return c.Error(err)
    }
    if err := i.repo.Delete(input.Images); err != nil {
        return c.Error(err)
    }
    c.String(http.StatusOK, http.StatusText(http.StatusOK))

    return nil
}
