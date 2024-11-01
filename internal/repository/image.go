package repository

import (
    "Scrapper/internal/entity"
    "Scrapper/internal/util"
    "gorm.io/gorm"
    "os"
    "path"
)

type Image struct {
    db *gorm.DB
    util.Config
}

func ImageRepository(db *gorm.DB, config util.Config) Image {
    return Image{db, config}
}

func (r *Image) SaveImage(scrapId int, fileNames []string) error {
    images := make([]entity.Image, len(fileNames))
    for idx, v := range fileNames {
        images[idx].FileName = v
        images[idx].ScrapID = scrapId
    }
    result := r.db.CreateInBatches(images, len(fileNames))
    return result.Error
}

// delete Fetched Image
func (r *Image) delete(images []entity.Image, tx ...*gorm.DB) error {
    var con *gorm.DB
    if len(tx) != 0 {
        con = tx[0]
    } else {
        con = r.db
    }

    con = con.Delete(&images)
    if con.Error != nil {
        for _, v := range images {
            file := path.Join(r.Media, v.FileName)
            if _, err := os.Stat(file); err != nil {
                _ = os.Remove(file)
            }
        }
    }
    return con.Error
}

func (r *Image) Delete(imageId []int) error {

    var images []entity.Image

    tx := r.db.Find(&images, imageId)
    err := r.delete(images, tx)
    return err
}

func (r *Image) DeleteByScrapId(scrapId int) error {
    images, err := r.getByScrapId(scrapId)
    if err != nil {
        err = r.delete(images)
    }
    return err
}
func (r *Image) getByScrapId(scrapId int) ([]entity.Image, error) {
    var images []entity.Image

    tx := r.db.Where(&entity.Image{ScrapID: scrapId}).Find(&images)
    return images, tx.Error
}

func (r *Image) ExportAndDelete(imageId []int) error {
    var images []entity.Image

    tx := r.db.Find(&images, imageId)
    tx = tx.Delete(&images)
    if tx.Error != nil {
        for _, v := range images {
            source := path.Join(r.Media, v.FileName)
            dest := path.Join(r.Export, v.FileName)
            if _, err := os.Stat(source); err != nil {
                _ = os.Rename(source, dest)
            }
        }
    }
    return tx.Error
}
