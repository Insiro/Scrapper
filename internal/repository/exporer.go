package repository

import (
	"Scrapper/internal/model/dto"
	"Scrapper/internal/model/entity"
	"gorm.io/gorm"
)

type ExporterRepository struct {
	db *gorm.DB
}

func (r *ExporterRepository) Create(data dto.CreateExporter) (entity.Exporter, error) {
	exp := entity.Exporter{Title: data.Title, Mode: data.Mode, NameRule: data.NameRule}
	result := r.db.Create(exp)
	return exp, result.Error
}

func (r *ExporterRepository) Get(exporterId uint, title string) ([]entity.Exporter, error) {
	var exp []entity.Exporter
	result := r.db.Where(&entity.Exporter{ID: exporterId, Title: title}).Find(exp)
	return exp, result.Error
}

func (r *ExporterRepository) Update(exporter entity.Exporter, data dto.UpdateExporter) (entity.Exporter, error) {
	if data.Title != nil {
		exporter.Title = *data.Title
	}
	if data.Mode != nil {
		exporter.Mode = *data.Mode
	}
	if data.NameRule != nil {
		exporter.NameRule = *data.NameRule
	}

	result := r.db.Save(exporter)
	return exporter, result.Error
}

func (r *ExporterRepository) Delete(exporterId uint, title string) error {
	result := r.db.Where(&entity.Exporter{ID: exporterId, Title: title}).Delete(&entity.Exporter{})
	return result.Error
}
