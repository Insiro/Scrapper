package entity

type Image struct {
    ID       uint   `gorm:"primaryKey;autoIncrement;index"`
    FileName string `gorm:"size:255"`
    ScrapID  uint   `gorm:"index"`
    Scrap    Scrap  `gorm:"foreignKey:ScrapID"`
}

func (Image) TableName() string {
    return "images"
}
