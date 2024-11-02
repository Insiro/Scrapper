package util

import (
	"Scrapper/internal/app"
	"image"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
)

func DownloadImage(url, fileName string, config *app.Config) (string, error) {
	// 1. URL에서 이미지 다운로드
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	imgConfig, format, err := image.DecodeConfig(res.Body)
	if err != nil {
		return "", err
	}

	// 3. 크기 필터링
	if imgConfig.Width < config.ImageWithThr || imgConfig.Height < config.ImageHeightThr {
		return "", nil
	}

	// 2. Content-Type으로 확장자 결정
	ext := ".jpg" // 기본 확장자
	if exts, _ := mime.ExtensionsByType("image/" + format); len(exts) > 0 {
		ext = exts[0]
	}

	// 3. 파일 생성 및 이미지 저장
	fname := fileName + ext
	fname = path.Join(path.Dir(config.Media), fname)
	file, err := os.Create(fname)
	if err != nil {
		return "", err
	}
	defer file.Close()

	if _, err = io.Copy(file, res.Body); err != nil {
		return "", err
	}

	return fname, nil
}
