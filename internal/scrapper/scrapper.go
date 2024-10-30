package scrapper

import (
	"Scrapper/internal/dto"
	"Scrapper/internal/entity/enum"
	"net/url"
	"strings"
)

type ScrapArgs struct {
	Url string
	Key string
}

type Scrapper interface {
	PreprocessURL(rawURL string) (*url.URL, error)
	GenArgs(parsedURL *url.URL) ScrapArgs
	MergeURL(parsedURL *url.URL) string
	createScrap(authorTag, authorName, sourceKey, content string, imageNames []string) dto.ScrapCreate
	Scrap(args *ScrapArgs) (dto.ScrapCreate, error)
}

type AbsScrapper struct {
	PageType enum.PageType
}

func (a *AbsScrapper) PreprocessURL(rawURL string) (*url.URL, error) {
	parsedURL, _ := url.Parse(rawURL)
	parsedURL.RawQuery = ""
	return parsedURL, nil
}

func (a *AbsScrapper) GenArgs(parsedURL *url.URL) ScrapArgs {
	return ScrapArgs{Url: a.MergeURL(parsedURL)}
}

func (a *AbsScrapper) MergeURL(parsedURL *url.URL) string {
	return parsedURL.String()
}

func (a *AbsScrapper) createScrap(authorTag, authorName, sourceKey, content string, imageNames []string) dto.ScrapCreate {
	// author 정보 처리
	if strings.HasPrefix(authorName, "@") {
		authorName = authorName[1:]
	}
	if strings.HasPrefix(authorTag, "@") {
		authorTag = authorTag[1:]
	}

	// 태그 필터링
	words := strings.Fields(content)
	var tags []string
	for _, word := range words {
		if strings.HasPrefix(word, "#") {
			tags = append(tags, word[1:])
		}
	}

	return dto.ScrapCreate{
		ScrapModifier: dto.ScrapModifier{
			AuthorName: authorName,
			AuthorTag:  authorTag,
			Content:    &content,
			Tags:       &tags,
		},
		SourceKey:  sourceKey,
		Source:     a.PageType,
		ImageNames: imageNames,
	}
}
