package scrapper

import (
    "Scrapper/internal/app"
    "Scrapper/internal/dto"
    "Scrapper/internal/scrapper/util"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/proto"
    "github.com/google/uuid"
    "net/url"
    "strings"
)

type implTwitter struct {
    AbsScrapper
}

var _ Scrapper = (*implTwitter)(nil)

const resourcePath = "https://pbs.twimg.com/media/*"

func filterTwitterRequest(ctx *rod.Hijack) {
    ext := "." + ctx.Request.URL().Query().Get("format")
    contain := app.IMAGE_EXTENSIONS.Contains(ext)

    if contain {
        ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
    }
    ctx.ContinueRequest(&proto.FetchContinueRequest{})
}

func (h *implTwitter) GenArgs(parsedURL *url.URL) ScrapArgs {
    args := h.AbsScrapper.GenArgs(parsedURL)
    path := parsedURL.Path
    if path[0] == '/' {
        path = path[1:]
    }
    if path[len(path)-1] == '/' {
        path = path[:len(path)-1]
    }
    args.Key = path
    return args
}
func (h *implTwitter) Scrap(args *ScrapArgs) (dto.ScrapCreate, error) {
    loader := util.NewContentLoader(nil).AddFilterRule(resourcePath, filterTwitterRequest)
    docText, err := loader.LoadHTMLContent(args.Url)
    if err != nil {
        return dto.ScrapCreate{}, nil
    }
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(docText))
    article := doc.Find("article").First()

    // author 찾기
    author := article.Find("div[data-testid='User-Name']").Find("a")
    name := author.Eq(0).Text()
    tag := author.Eq(1).Text()

    wrapper := article.Children().Filter("div").Children().Filter("div").Children().Filter("div").Eq(2)

    // content 찾기 및 텍스트 추출
    content := wrapper.Find("div").Find("div").Find("div")
    contentTxt := content.Text()

    // img_list 찾기
    imgList := content.Find("img")
    var fnameList []string

    // 이미지 다운로드 및 fnameList 추가
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
    return h.createScrap(tag, name, args.Key, contentTxt, fnameList), nil
}
