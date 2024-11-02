package util

import (
    "Scrapper/internal/app"
    "context"
    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/launcher"
    "github.com/go-rod/rod/lib/proto"
    "time"

    "github.com/chromedp/chromedp"
)

type hiject struct {
    Pattern string
    Handler func(*rod.Hijack)
}
type LoaderConfig struct {
    Headless      *bool
    FilterRequest *bool
    FilterImage   *bool
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
        if config.FilterImage != nil {
            filterImage = *config.FilterImage
        }
        if config.Headless != nil {
            headless = *config.Headless
        }
        if config.FilterRequest != nil {
            filterRequest = *config.FilterRequest
        }
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

    content, _ := page.MustElement(`body`).HTML()

    return content, nil
}

// LoadHTMLContent는 페이지 HTML을 로드하고 특정 요청을 차단합니다.
func LoadHTMLContent2(url string) (string, error) {
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
        chromedp.Flag("Headless", false),    // 헤드리스 모드를 끔
        chromedp.Flag("disable-gpu", false), // GPU 사용 가능
        chromedp.Flag("no-sandbox", true),   // 샌드박스 비활성화 (일부 시스템에서 필요할 수 있음)
        chromedp.Flag("disable-javascript", false),
        chromedp.WindowSize(1280, 1024), // 창 크기 설정
    )
    // remote 디버깅 주소에 연결
    // ExecAllocator 생성
    allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
    defer cancel()
    //    allocCtx, cancel := chromedp.NewRemoteAllocator(context.Background(), "http://localhost:9222")
    //    defer cancel()

    // chromedp 컨텍스트 생성
    ctx, cancel := chromedp.NewContext(allocCtx)
    defer cancel()

    // 뷰포트 크기 설정 등 옵션
    chromedp.Run(ctx, chromedp.EmulateViewport(2560, 1440))

    var htmlContent string

    err := chromedp.Run(ctx,
        chromedp.Navigate(url), // 페이지 탐색
        //        chromedp.WaitReady("root"), // 페이지의 body 요소가 로드될 때까지 대기
        //        chromedp.Sleep(5*time.Second),

        chromedp.WaitReady("body"),
    )
    time.Sleep(time.Second * 2)
    if err != nil {
        return "", err
    }
    err = chromedp.Run(ctx,
        chromedp.ActionFunc(func(ctx context.Context) error {
            for {
                var readyState string
                err := chromedp.Evaluate(`document.readyState`, &readyState).Do(ctx)
                if err != nil {
                    return err
                }
                if readyState == "complete" {
                    break
                }
                time.Sleep(500 * time.Millisecond) // 0.5초 간격으로 계속 확인
            }
            return nil
        }),
        chromedp.OuterHTML("html", &htmlContent),
    )

    if err != nil {
        return "", err
    }
    return htmlContent, nil
}
