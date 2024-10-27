package util

import (
	"io"
	"mime"
	"net/http"
	"os"
)

func DownloadImage(url, fileName string) (string, error) {
	// 1. URL에서 이미지 다운로드
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// 2. Content-Type으로 확장자 결정
	ext := ".jpg" // 기본 확장자
	if exts, _ := mime.ExtensionsByType(res.Header.Get("Content-Type")); len(exts) > 0 {
		ext = exts[0]
	}

	// 3. 파일 생성 및 이미지 저장
	fname := fileName + ext
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
