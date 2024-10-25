package model

type Tag struct {
	Name    string `gorm:"size:200;primaryKey;index"`
	ScrapID uint   `gorm:"primaryKey;index"`
}

func (Tag) TableName() string {
	return "tags"
}
