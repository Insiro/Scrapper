package controller

import (
	"Scrapper/internal/model/dto"
	"Scrapper/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExporterController struct {
	repo repository.ExporterRepository
}

func (r *ExporterController) Create(c *gin.Context) *gin.Error {
	var input = dto.CreateExporter{}
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

func (r *ExporterController) Update(c *gin.Context) *gin.Error {

	var input = dto.UpdateExporter{}
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

func (r *ExporterController) List(c *gin.Context) *gin.Error {
	var input = dto.SelectExporter{}
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

func (r *ExporterController) Init(group *gin.RouterGroup) {

}
