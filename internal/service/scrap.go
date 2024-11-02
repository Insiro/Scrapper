package service

import (
    "Scrapper/internal/app"
    "Scrapper/internal/appError"
    "Scrapper/internal/dto"
    "Scrapper/internal/entity"
    "Scrapper/internal/repository"
    "Scrapper/internal/scrapper"
)

type Scrap struct {
    repo    repository.Scrap
    imgRepo repository.Image
    config  *app.Config
}

func ScrapService(repo repository.Scrap, imgRepo repository.Image, config *app.Config) Scrap {
    return Scrap{repo, imgRepo, config}
}

func (s *Scrap) Scrap(target string) (*entity.Scrap, error) {
    scraper, err := scrapper.Factory(target, nil, nil, s.config)
    if err != nil {
        return nil, err
    }
    args := scraper.Args
    if exist, err := s.repo.GetBySourceId(scraper.PageType, args.Key); err == nil {
        return &exist, appError.Duplicated{}
    }

    scrapCreate, err := scraper.Scrap(nil)
    if err != nil {
        return nil, err
    }
    saved, err := s.repo.Create(scrapCreate)
    if err != nil {
        return nil, err
    }
    err = s.imgRepo.SaveImage(saved.ID, scrapCreate.ImageNames)
    if err != nil {
        return nil, err
    }
    return &saved, nil
}

func (s *Scrap) ReScrap(id int) (*entity.Scrap, error) {
    found, err := s.repo.GetScrap(id)
    if err != nil {
        return nil, err
    }
    scr, _ := scrapper.Factory(found.Url(), &found.Source, nil, s.config)
    data, err := scr.Scrap(nil)
    if err != nil {

        return nil, err
    }
    updateData := dto.Create2Update(data)

    result, err := s.repo.Update(found, updateData)
    if err != nil {
        return nil, err
    }
    err = s.imgRepo.DeleteByScrapId(id)
    if err != nil {
        return nil, err
    }

    err = s.imgRepo.SaveImage(id, data.ImageNames)
    if err != nil {
        return nil, err
    }

    return &result, nil
}

func (s *Scrap) Get(id int) (*entity.Scrap, error) {
    scrap, err := s.repo.GetScrap(id)
    if err != nil {
        return nil, err
    }
    //TODO: load Image
    tags, _ := s.repo.GetTags(id, "")
    scrap.Tags = tags
    return &scrap, nil
}

func (s *Scrap) Count() (int64, error) {
    count, err := s.repo.CountScrap()
    return count, err
}

func (s *Scrap) List(offset int, limit int, pined bool) ([]entity.Scrap, error) {
    scraps, err := s.repo.ListScrap(offset, limit, pined)
    for _, scrap := range scraps {
        tag, _ := s.repo.GetTags(scrap.ID, "")
        scrap.Tags = tag
    }
    return scraps, err
}

func (s *Scrap) Update(id int, update dto.ScrapUpdate) (*entity.Scrap, error) {
    scrap, err := s.repo.GetScrap(id)
    if err != nil {
        return nil, err
    }
    scrap, err = s.repo.Update(scrap, update)
    if err != nil {
        return nil, err
    }
    tags, _ := s.repo.GetTags(id, "")
    scrap.Tags = tags

    return &scrap, nil
}

func (s *Scrap) Delete(id int) error {
    if err := s.repo.DeleteScrap(id); err != nil {
        return err
    }

    if err := s.imgRepo.DeleteByScrapId(id); err != nil {
        return err
    }

    return nil
}
