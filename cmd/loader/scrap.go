package loader

import (
	"Scrapper/internal/app"
	"Scrapper/internal/repository"
	"Scrapper/internal/service"
	"Scrapper/pkg/out"
	"gorm.io/gorm"
	"os"
)

func Scrap(target string, db *gorm.DB, config *app.Config) {
	scService := service.ScrapService(repository.ScrapRepository(db), repository.ImageRepository(db, config), config)
	scrap, err := scService.Scrap(target)
	if err != nil {
		panic("failed to scrap")
	}

	out.Table(scrap)
	os.Exit(1)
}
