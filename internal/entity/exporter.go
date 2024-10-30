package entity

import "Scrapper/internal/entity/enum"

type Exporter struct {
    ID       int             `gorm:"primaryKey;autoIncrement;index" json:"id,omitempty"`
    Title    string          `gorm:"size:255" json:"title,omitempty"`
    Mode     enum.ExportMode `gorm:"type:enum('FULL', 'IMG', 'TXT')" json:"mode,omitempty"`
    NameRule string          `gorm:"size:255; default:''" json:"name_rule,omitempty"`
}

func (Exporter) TableName() string {
    return "exporter"
}
