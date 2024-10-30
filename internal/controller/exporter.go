package controller

import (
    "Scrapper/internal/dto"
    "Scrapper/internal/repository"
    "github.com/gin-gonic/gin"
    "net/http"
)

type Exporter struct {
    repo repository.Exporter
}

func (r *Exporter) Create(c *gin.Context) *gin.Error {
    var input = dto.ExporterCreate{}
    if err := c.ShouldBindJSON(&input); err != nil {
        return c.Error(err)
    }
    exporter, err := r.repo.Create(input)
    if err != nil {
        return c.Error(err)
    }

    c.JSON(http.StatusCreated, exporter)
    return nil
}

func (r *Exporter) Update(c *gin.Context) *gin.Error {

    var input = dto.ExporterUpdate{}
    if err := c.ShouldBindJSON(&input); err != nil {
        return c.Error(err)
    }
    found, err := r.repo.Get(input.Id, "")
    if err != nil {
        return c.Error(err)
    }
    exporter, err := r.repo.Update(found[0], input)
    if err != nil {
        return c.Error(err)
    }

    c.JSON(http.StatusOK, exporter)
    return nil
}

func (r *Exporter) List(c *gin.Context) *gin.Error {
    var input = dto.ExporterSelect{}
    if err := c.ShouldBindJSON(&input); err != nil {
        return c.Error(err)
    }
    found, err := r.repo.Get(input.Id, input.Title)
    if err != nil {
        return c.Error(err)
    }

    c.JSON(http.StatusOK, found)
    return nil
}

func (r *Exporter) Init(group *gin.RouterGroup) Exporter {
    return *r
}
