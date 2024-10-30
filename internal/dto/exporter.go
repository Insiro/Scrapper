package dto

import "Scrapper/internal/entity/enum"

type ExporterCreate struct {
    Title    string
    Mode     enum.ExportMode
    NameRule string
}
type ExporterUpdate struct {
    Id       int
    Title    *string
    Mode     *enum.ExportMode
    NameRule *string
}
type ExporterSelect struct {
    Id    int
    Title string
}
