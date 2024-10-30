package repository

import (
    "Scrapper/internal/dto"
    "Scrapper/internal/entity"
    "gorm.io/gorm"
)

type Exporter struct {
    db *gorm.DB
}

func (r *Exporter) Create(data dto.ExporterCreate) (entity.Exporter, error) {
    exp := entity.Exporter{Title: data.Title, Mode: data.Mode, NameRule: data.NameRule}
    result := r.db.Create(exp)
    return exp, result.Error
}

func (r *Exporter) Get(exporterId int, title string) ([]entity.Exporter, error) {
    var exp []entity.Exporter
    result := r.db.Where(&entity.Exporter{ID: exporterId, Title: title}).Find(exp)
    return exp, result.Error
}

func (r *Exporter) Update(exporter entity.Exporter, data dto.ExporterUpdate) (entity.Exporter, error) {
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

func (r *Exporter) Delete(exporterId int, title string) error {
    result := r.db.Where(&entity.Exporter{ID: exporterId, Title: title}).Delete(&entity.Exporter{})
    return result.Error
}
