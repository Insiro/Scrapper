package loader

import (
    "Scrapper/internal/model/entity"
    "fmt"
    "gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
    if err := db.AutoMigrate(&entity.Scrap{}, &entity.Image{}, &entity.Tag{}); err != nil {
        fmt.Println("Migration failed")
        fmt.Println(err.Error())
        panic(err)
    }
    fmt.Println("Migration successful")

}
