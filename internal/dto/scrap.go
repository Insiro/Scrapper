package dto

import (
    "Scrapper/internal/entity"
    "Scrapper/internal/entity/enum"
)

type URLInput struct {
    Url string
}

type ListScrap struct {
    Offset int  `form:"offset,default=0"`
    Limit  int  `form:"offset,default=20"`
    Pined  bool `form:"pined,default=false"`
}

type ScrapModifier struct {
    Content    *string
    AuthorName string
    AuthorTag  string
    Pin        *bool
    Comment    *string
    Tags       *[]string
}

type ScrapCreate struct {
    ScrapModifier
    SourceKey  string
    Source     enum.PageType
    ImageNames []string
}

type ScrapUpdate struct {
    ScrapModifier
    DeleteImages []int
    NewImages    []int
}

func Create2Update(create ScrapCreate) ScrapUpdate {
    return ScrapUpdate{
        ScrapModifier: ScrapModifier{
            Content:    create.Content,
            AuthorName: create.AuthorName,
            AuthorTag:  create.AuthorTag,
            Pin:        create.Pin,
            Comment:    create.Comment,
            Tags:       create.Tags,
        },
    }
}

type Scrap struct {
    Id         int
    Source     enum.PageType
    SourceId   string
    Url        string
    Content    *string
    AuthorName string
    AuthorTag  string
    Comment    *string
    Images     []Image
    Pin        bool
    Tags       []string
}

func NewScrap(scrap entity.Scrap, tags []entity.Tag) Scrap {
    tagNames := make([]string, len(tags))
    for i, tag := range tags {
        tagNames[i] = tag.Name
    }

    imgList := make([]Image, len(scrap.Images))
    for i, img := range scrap.Images {
        imgList[i] = Image{img.ID, img.FileName}
    }

    return Scrap{
        Id:         scrap.ID,
        SourceId:   scrap.SourceID,
        Url:        scrap.Source.Url(scrap.SourceID),
        Content:    scrap.Content,
        AuthorName: scrap.AuthorName,
        AuthorTag:  scrap.AuthorTag,
        Comment:    scrap.Comment,
        Images:     imgList,
        Source:     scrap.Source,
        Pin:        scrap.Pin,
        Tags:       tagNames,
    }
}
