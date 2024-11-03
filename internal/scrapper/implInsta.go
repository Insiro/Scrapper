package scrapper

import (
    "Scrapper/internal/dto"
    "Scrapper/internal/scrapper/util"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/google/uuid"
    "net/url"
    "strings"
)

type implInsta struct {
    AbsScrapper
}

var _ Scrapper = (*implInsta)(nil)

func (h *implInsta) GenArgs(parsedURL *url.URL) ScrapArgs {
    args := h.AbsScrapper.GenArgs(parsedURL)

    part := strings.Split(parsedURL.Path, "/p/")
    part = strings.Split(part[1], "/")
    args.Key = part[0]

    return args
}

func (h *implInsta) Scrap(args *ScrapArgs) (dto.ScrapCreate, error) {
    loader := util.NewContentLoader(&util.LoaderConfig{FilterRequest: true, FilterImage: true})
    docText, err := loader.LoadHTMLContent(args.Url)
    if err != nil {
        return dto.ScrapCreate{}, nil
    }
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(docText))

    userId := ""
    doc.Find("meta").Each(func(i int, s *goquery.Selection) {
        if property, _ := s.Attr("property"); property == "instapp:owner_user_id" {
            userId, _ = s.Attr("content")
        }
    })

    wrapper := doc.Find("article").ChildrenFiltered("div").ChildrenFiltered("div")
    wrapperFirst := wrapper.First()
    imgList := wrapperFirst.Find("img")
    wrapper2 := wrapperFirst.Next()
    authorTag := wrapper2.Find("header").ChildrenFiltered("div").First().Next().ChildrenFiltered("div")
    name := authorTag.ChildrenFiltered("div")
    // author 찾기

    section := wrapper2.Find("section")
    content := section.Next().Next().Find("h2").Next()

    // 이미지 다운로드 및 fnameList 추가
    var fnameList []string
    imgList.Each(func(i int, s *goquery.Selection) {
        src, exists := s.Attr("src")
        if exists {
            // UUID로 고유한 파일 이름 생성
            filename := uuid.New().String()
            fmt.Println(src, filename)
            fname, err := util.DownloadImage(src, filename, h.config)
            if err == nil {
                fnameList = append(fnameList, fname)
            }
        }
    })
    return h.createScrap(userId, name.Text(), args.Key, content.Text(), fnameList), nil
}
