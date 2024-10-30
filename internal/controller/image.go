package controller

import (
    "Scrapper/internal/dto"
    "Scrapper/internal/repository"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type Image struct {
    repo repository.Image
}

func (i *Image) Delete(c *gin.Context) *gin.Error {
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

func (i *Image) DeleteList(c *gin.Context) *gin.Error {
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
