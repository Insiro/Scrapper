package repository

import (
	"Scrapper/internal/model/entity"
	"gorm.io/gorm"
	"os"
	"path"
)

type ImageRepository struct {
	db         *gorm.DB
	mediaPath  string
	exportPath string
}

func (r *ImageRepository) SaveImage(scrapId uint, fileNames []string) error {
	images := make([]entity.Image, len(fileNames))
	for idx, v := range fileNames {
		images[idx].FileName = v
		images[idx].ScrapID = scrapId
	}
	result := r.db.CreateInBatches(images, len(fileNames))
	return result.Error
}

// delete Fetched Images
func (r *ImageRepository) delete(images []entity.Image, tx ...*gorm.DB) error {
	var con *gorm.DB
	if len(tx) != 0 {
		con = tx[0]
	} else {
		con = r.db
	}

	con = con.Delete(&images)
	if con.Error != nil {
		for _, v := range images {
			file := path.Join(r.mediaPath, v.FileName)
			if _, err := os.Stat(file); err != nil {
				_ = os.Remove(file)
			}
		}
	}
	return con.Error
}

func (r *ImageRepository) Delete(imageId []uint) error {

	var images []entity.Image

	tx := r.db.Find(&images, imageId)
	err := r.delete(images, tx)
	return err
}

func (r *ImageRepository) DeleteByScrapId(scrapId uint) error {
	images, err := r.getByScrapId(scrapId)
	if err != nil {
		err = r.delete(images)
	}
	return err
}
func (r *ImageRepository) getByScrapId(scrapId uint) ([]entity.Image, error) {
	var images []entity.Image

	tx := r.db.Where(&entity.Image{ScrapID: scrapId}).Find(&images)
	return images, tx.Error
}

func (r *ImageRepository) ExportAndDelete(imageId []uint) error {
	var images []entity.Image

	tx := r.db.Find(&images, imageId)
	tx = tx.Delete(&images)
	if tx.Error != nil {
		for _, v := range images {
			source := path.Join(r.mediaPath, v.FileName)
			dest := path.Join(r.exportPath, v.FileName)
			if _, err := os.Stat(source); err != nil {
				_ = os.Rename(source, dest)
			}
		}
	}
	return tx.Error
}
