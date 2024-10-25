package model

type Scrap struct {
	ID         uint    `gorm:"primaryKey;autoIncrement;index"`
	Pin        bool    `gorm:"default:false;index"`
	SourceID   string  `gorm:"size:255;uniqueIndex:scraps_unique"`
	Source     string  `gorm:"size:50;uniqueIndex:scraps_unique"`
	Content    *string `gorm:"type:text"`
	AuthorName string  `gorm:"size:100"`
	AuthorTag  string  `gorm:"size:100"`
	Comment    *string `gorm:"size:255"`
}

func (Scrap) TableName() string {
	return "scraps"
}
