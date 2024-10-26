package loader

import (
    "Scrapper/internal/model"
    "fmt"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    if err := db.AutoMigrate(&model.Scrap{}, &model.Image{}, &model.Exporter{}, &model.Tag{}); err != nil {
        fmt.Println("Migration failed")
        panic(err)
    }
    fmt.Println("Migration successful")

}
