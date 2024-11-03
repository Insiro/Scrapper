package util

import (
    "Scrapper/internal/app"
    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/launcher"
    "github.com/go-rod/rod/lib/proto"
)

type hiject struct {
    Pattern string
    Handler func(*rod.Hijack)
}
type LoaderConfig struct {
    Headless      bool
    FilterRequest bool
    FilterImage   bool
}

type contentLoader struct {
    Headless      bool
    FilterRequest bool
    FilterImage   bool
    hiject        []hiject
}

func NewContentLoader(config *LoaderConfig) *contentLoader {
    headless := true
    filterImage := true
    filterRequest := true
    if config != nil {
        headless = config.Headless
        filterImage = config.FilterImage
        filterRequest = config.FilterRequest
    }

    return &contentLoader{Headless: headless, FilterImage: filterImage, FilterRequest: filterRequest, hiject: make([]hiject, 0)}
}

func (s *contentLoader) AddFilterRule(pattern string, handler func(*rod.Hijack)) *contentLoader {
    s.hiject = append(s.hiject, hiject{Pattern: pattern, Handler: handler})
    return s
}

func (s *contentLoader) LoadHTMLContent(url string) (string, error) {
    u := launcher.New().Headless(s.Headless).MustLaunch()
    browser := rod.New().ControlURL(u).MustConnect()
    defer browser.MustClose()

    if s.FilterRequest {
        router := browser.HijackRequests()
        defer router.MustStop()

        if s.FilterImage {
            for ext := range app.IMAGE_EXTENSIONS {
                router.MustAdd("*"+ext, func(ctx *rod.Hijack) {
                    if ctx.Request.Type() == proto.NetworkResourceTypeImage {
                        ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
                        return
                    }
                    ctx.ContinueRequest(&proto.FetchContinueRequest{})
                })
            }
        }
        for _, item := range s.hiject {
            router.MustAdd(item.Pattern, item.Handler)
        }

        go router.Run()
    }

    page := browser.MustPage(url)

    page.MustWaitLoad() // 페이지 로드가 완료될 때까지 대기
    page.MustWaitIdle() // 네트워크 안정화 대기
    page.MustWaitRequestIdle()
    page.MustWaitDOMStable()

    content, _ := page.MustElement(`html`).HTML()

    return content, nil
}
