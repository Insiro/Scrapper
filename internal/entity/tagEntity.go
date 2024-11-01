package entity

type Tag struct {
    Id      int    `gorm:"primaryKey;autoIncrement;index"`
    Name    string `gorm:"size:200;uniqueIndex"`
    ScrapID int    `gorm:"uniqueIndex"`
}

func (Tag) TableName() string {
    return "tags"
}
