package scrapper

import (
	"Scrapper/internal/model/dto"
	"Scrapper/internal/model/entity/pageType"
	"Scrapper/internal/util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type implHoyolab struct {
	AbsScrapper
}

func NewImplHoyolab() *implHoyolab {
	return &implHoyolab{AbsScrapper{PageType: pageType.HoyoLab}}
}

var _ Scrapper = (*implHoyolab)(nil)

func (h *implHoyolab) GenArgs(parsedURL *url.URL) ScrapArgs {
	args := h.AbsScrapper.GenArgs(parsedURL)
	args.Key = strings.Split(parsedURL.Path, "article")[1][1:]
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

	client := &http.Client{}
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

	imgList := content["imgs"].([]string)
	contentText := content["describe"].(string)

	var fnameList []string
	for _, img := range imgList {
		uid := uuid.New()
		fname, err := util.DownloadImage(img, uid.String())
		if err == nil {
			fnameList = append(fnameList, fmt.Sprintf(fname))
		}
	}

	return h.createScrap(userData["uid"].(string), userData["nickname"].(string), args.Key, title+"\n"+contentText, fnameList), nil
}

type implHoyolink struct {
	implHoyolab
}

func (h *implHoyolink) PreprocessUrl(url string) (*url.URL, error) {
	// HTTP GET 요청을 보내고 User-Agent 헤더 설정
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", util.USER_AGENT)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// goquery로 HTML 파싱
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	// 모든 meta 태그를 찾아 리디렉션 URL을 추출
	var redirectURL string
	doc.Find("meta").EachWithBreak(func(i int, s *goquery.Selection) bool {
		content, exists := s.Attr("content")
		if exists && strings.HasPrefix(content, "0;url=") {
			redirectURL = content[7:] // URL 부분만 추출
			return false              // break 루프
		}
		return true // continue 루프
	})

	if redirectURL == "" {
		return nil, fmt.Errorf("no redirect URL found")
	}

	return h.implHoyolab.PreprocessURL(redirectURL)
}
