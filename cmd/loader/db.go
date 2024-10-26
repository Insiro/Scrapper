package loader

import (
    "Scrapper/internal/model/entity"
    "fmt"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    if err := db.AutoMigrate(&entity.Scrap{}, &entity.Image{}, &entity.Exporter{}, &entity.Tag{}); err != nil {
        fmt.Println("Migration failed")
        panic(err)
    }
    fmt.Println("Migration successful")

}
