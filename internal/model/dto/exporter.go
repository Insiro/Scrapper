package dto

import "Scrapper/internal/model/entity"

type CreateExporter struct {
	Title    string
	Mode     entity.ExportMode
	NameRule string
}
type UpdateExporter struct {
	Id       int
	Title    *string
	Mode     *entity.ExportMode
	NameRule *string
}
type SelectExporter struct {
	Id    *int
	Title *string
}
