package scrapper

import (
    "Scrapper/internal/dto"
    "Scrapper/internal/scrapper/util"
    "errors"
    "fmt"
    "github.com/goccy/go-json"
    "github.com/google/uuid"
    "io"
    "net/http"
    "net/url"
    "os/exec"
    "strings"
)

type implHoyolab struct {
    AbsScrapper
}

var _ Scrapper = (*implHoyolab)(nil)

func (h *implHoyolab) GenArgs(parsedURL *url.URL) ScrapArgs {
    args := h.AbsScrapper.GenArgs(parsedURL)
    var key string
    if parsedURL.Path == "/" {
        key = strings.Split(parsedURL.Fragment, "article/")[1]
    } else {
        key = strings.Split(parsedURL.Fragment, "article/")[1]
    }
    args.Key = strings.Split(key, "/")[0]

    return args
}

func (h *implHoyolab) Scrap(args *ScrapArgs) (dto.ScrapCreate, error) {
    apiPath := fmt.Sprintf("https://bbs-api-os.hoyolab.com/community/post/wapi/getPostFull?post_id=%s", args.Key)
    req, err := http.NewRequest("GET", apiPath, nil)
    if err != nil {
        return dto.ScrapCreate{}, err
    }
    req.Header.Set("x-rpc-language", "ko-kr")
    req.Header.Set("x-rpc-timezone", "Asia/Seoul")
    req.Header.Set("x-rpc-show-translated", "true")

    client := &http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }}
    resp, err := client.Do(req)
    if err != nil {
        return dto.ScrapCreate{}, err
    }
    defer resp.Body.Close()

    var result map[string]interface{}
    body, _ := io.ReadAll(resp.Body)
    err = json.Unmarshal(body, &result)
    if err != nil {
        return dto.ScrapCreate{}, err
    }

    post := result["data"].(map[string]interface{})["post"].(map[string]interface{})
    userData := post["user"].(map[string]interface{})
    postData := post["post"].(map[string]interface{})

    title := postData["subject"].(string)
    contentData := postData["content"].(string)
    var content map[string]interface{}
    err = json.Unmarshal([]byte(contentData), &content)
    if err != nil {
        return dto.ScrapCreate{}, err
    }

    imgList := content["imgs"].([]interface{})

    contentText := content["describe"].(string)

    var fnameList []string
    for _, img := range imgList {
        uid := uuid.New()
        fname, err := util.DownloadImage(fmt.Sprint(img), uid.String(), h.config)
        if err == nil {
            fnameList = append(fnameList, fmt.Sprintf(fname))
        }
    }

    return h.createScrap(userData["uid"].(string), userData["nickname"].(string), args.Key, title+"\n"+contentText, fnameList), nil
}

type implHoyolink struct {
    implHoyolab
}

var _ Scrapper = (*implHoyolink)(nil)

func (h implHoyolink) PreprocessURL(urlPath string) (*url.URL, error) {
    cmd := exec.Command("curl", "-s", urlPath)
    output, err := cmd.Output()
    if err != nil {
        return nil, err
    }
    parts := strings.Split(string(output), "\"")
    if len(parts) <= 1 {
        return nil, errors.New("failed to Preprocess HoyoLink 1")
    }
    next := parts[1]

    cmd = exec.Command("curl", "-s", next)
    output, err = cmd.Output()
    parts = strings.Split(string(output), "\"")
    if len(parts) <= 1 {
        return nil, errors.New("failed to Preprocess HoyoLink 2")
    }
    next = parts[1]

    p := strings.Split(next, ";")[0]
    parsed, err := url.Parse(p)
    que, err := url.ParseQuery(parsed.RawQuery)
    next = que.Get("url")

    return h.implHoyolab.PreprocessURL(next)
}
