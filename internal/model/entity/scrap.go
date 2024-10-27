package entity

import "Scrapper/internal/model/entity/pageType"

type Scrap struct {
    ID         uint              `gorm:"primaryKey;autoIncrement;index"`
    Pin        bool              `gorm:"default:false;index"`
    SourceID   string            `gorm:"size:255;uniqueIndex:scraps_unique"`
    Source     pageType.PageType `gorm:"size:50;uniqueIndex:scraps_unique"`
    Content    *string           `gorm:"type:text"`
    AuthorName string            `gorm:"size:100"`
    AuthorTag  string            `gorm:"size:100"`
    Comment    *string           `gorm:"size:255"`
    Images     []Image
}

func (Scrap) TableName() string {
    return "scraps"
}
func (sc *Scrap) Url() string {
    return sc.Source.Url(sc.SourceID)
}
