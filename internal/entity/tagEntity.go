package entity

type Tag struct {
    Name    string `gorm:"size:200;primaryKey;index"`
    ScrapID int    `gorm:"primaryKey;index"`
}

func (Tag) TableName() string {
    return "tags"
}
