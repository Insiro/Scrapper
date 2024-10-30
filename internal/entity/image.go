package entity

type Image struct {
    ID       int    `gorm:"primaryKey;autoIncrement;index"`
    FileName string `gorm:"size:255"`
    ScrapID  int    `gorm:"index"`
    Scrap    Scrap  `gorm:"foreignKey:ScrapID"`
}

func (Image) TableName() string {
    return "images"
}
