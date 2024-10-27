package dto

import (
    "Scrapper/internal/model/entity"
    "Scrapper/internal/model/entity/pageType"
)

type URLInput struct {
    Url string
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
    Source     pageType.PageType
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

type ScrapResponse struct {
    Id         uint
    Source     pageType.PageType
    SourceId   string
    Url        string
    Content    *string
    AuthorName string
    AuthorTag  string
    Comment    *string
    Images     []ImageResponse
    Pin        bool
    Tags       []string
}

func NewScrapResponse(scrap entity.Scrap, tags []entity.Tag) ScrapResponse {
    tagNames := make([]string, len(tags))
    for i, tag := range tags {
        tagNames[i] = tag.Name
    }

    images := make([]ImageResponse, len(scrap.Images))
    for i, img := range scrap.Images {
        images[i] = ImageResponse{img.ID, img.FileName}
    }

    return ScrapResponse{
        Id:         scrap.ID,
        SourceId:   scrap.SourceID,
        Url:        scrap.Source.Url(scrap.SourceID),
        Content:    scrap.Content,
        AuthorName: scrap.AuthorName,
        AuthorTag:  scrap.AuthorTag,
        Comment:    scrap.Comment,
        Images:     images,
        Source:     scrap.Source,
        Pin:        scrap.Pin,
        Tags:       tagNames,
    }
}
