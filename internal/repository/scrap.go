package repository

import (
    "Scrapper/internal/dto"
    "Scrapper/internal/entity"
    "Scrapper/internal/entity/enum"
    "gorm.io/gorm"
)

type Scrap struct {
    db *gorm.DB
}

func (r *Scrap) Create(scrapData dto.ScrapCreate) (entity.Scrap, error) {
    scrap := entity.Scrap{
        SourceID:   scrapData.SourceKey,
        AuthorName: scrapData.AuthorName,
        AuthorTag:  scrapData.AuthorTag,
        Source:     scrapData.Source,
        Content:    scrapData.Content,
        Comment:    scrapData.Comment,
    }
    tx := r.db.Begin()
    tx = tx.Create(scrap)
    if scrapData.Tags != nil {
        _ = r.PutTag(scrap.ID, *scrapData.Tags, tx)
    }
    tx = tx.Commit()
    if tx.Error != nil {
        tx.Rollback()
    }

    return scrap, tx.Error

}
func (r *Scrap) Update(scrap entity.Scrap, data dto.ScrapUpdate) (entity.Scrap, error) {
    if data.AuthorName != "" {
        scrap.AuthorName = data.AuthorName
    }
    if data.AuthorTag != "" {
        scrap.AuthorTag = data.AuthorTag
    }
    if data.Content != nil {
        scrap.Content = data.Content
    }
    if data.Comment != nil {
        scrap.Comment = data.Comment
    }
    if data.Pin != nil {
        scrap.Pin = *data.Pin
    }
    tx := r.db.Begin()
    tx = tx.Save(data)

    if data.Tags != nil {
        _ = r.PutTag(scrap.ID, *data.Tags, tx)
    }
    tx = tx.Commit()
    if tx.Error != nil {
        tx.Rollback()
    }

    return scrap, nil
}

func (r *Scrap) GetBySourceId(pageType enum.PageType, sourceId string) (entity.Scrap, error) {
    var scrap entity.Scrap
    result := r.db.Model(entity.Scrap{Source: pageType, SourceID: sourceId}).Take(&scrap)
    if result.Error != nil {
        return scrap, result.Error
    }
    return scrap, nil
}

func (r *Scrap) GetScrap(scrapId int) (entity.Scrap, error) {
    var scrap entity.Scrap
    result := r.db.Take(&scrap, scrapId)
    if result.Error != nil {
        return scrap, result.Error
    }
    return scrap, nil
}

func (r *Scrap) ListScrap(offset int, limit int, pined bool) ([]entity.Scrap, error) {

    q := r.db.Model(entity.Scrap{})
    if pined {
        q = q.Where(entity.Scrap{Pin: pined})
    }
    var scraps []entity.Scrap
    q = q.Order("id desc").Offset(offset).Limit(limit).Find(&scraps)
    return scraps, q.Error
}

func (r *Scrap) CountScrap() (int64, error) {
    var count int64
    result := r.db.Model(entity.Scrap{}).Count(&count)
    return count, result.Error
}

func (r *Scrap) DeleteScrap(scrapId int) error {
    tx := r.db.Begin()
    tx = tx.Delete(&Scrap{}, scrapId)
    tx = tx.Where("scrap_id = ?", scrapId).Delete(&entity.Image{})
    tx = tx.Commit()
    if tx.Error != nil {
        tx.Rollback()
    }
    return tx.Error
}

func (r *Scrap) PutTag(scrapId int, tagList []string, tx ...*gorm.DB) error {
    internal := len(tx) == 0
    var con *gorm.DB
    if internal {
        con = r.db.Begin()
    } else {
        con = tx[0]
    }
    tags := make([]entity.Tag, len(tagList))
    for i, v := range tagList {
        tags[i] = entity.Tag{Name: v, ScrapID: scrapId}
    }
    con = con.Delete(&entity.Tag{ScrapID: scrapId}).CreateInBatches(tags, len(tags))
    if internal {
        con = con.Commit()
        err := con.Error
        if con.Error != nil {
            con.Rollback()
            return err
        }
    }
    return nil
}

func (r *Scrap) GetTags(scrapId int, tagName string) ([]entity.Tag, error) {
    tx := r.db
    if scrapId != 0 {
        tx = tx.Where("scrap_id = ?", scrapId)
    }
    if tagName != "" {
        tx = tx.Where("tagName = ?", tagName)
    }
    var tags []entity.Tag
    result := tx.Find(&entity.Tag{}, tags)

    return tags, result.Error
}
