package app

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
)

type Config struct {
	Storage    string
	Media      string
	Export     string
	DBDriver   string
	DBName     string
	DBPassword string
	DBUsername string
	DBHost     string
	DBPort     string
	DBURL      string
	BaseURL    string

	ImageWithThr   int
	ImageHeightThr int
}

func InitConfig() *Config {
	// 기본 경로 설정
	storage := getEnv("SCRAPER_STORAGE", "./storage") + "/test"
	media := filepath.Join(storage, "media")
	export := filepath.Join(storage, "export")

	// 데이터베이스 설정
	dbDriver := getEnv("SCRAPER_DB_DRIVER", "sqlite")
	dbName := getEnv("SCRAPER_DB_NAME", "scrapper.db")
	dbPassword := getEnv("SCRAPER_DB_PASSWORD", "")
	dbUsername := getEnv("SCRAPER_DB_USERNAME", "")
	dbHost := getEnv("SCRAPER_DB_HOST", "")
	dbPort := getEnv("SCRAPER_DB_PORT", "5432")

	// 데이터베이스 URL 생성
	var dbURL string
	if dbDriver == "sqlite" {
		dbURL = filepath.Join(storage, dbName)
	} else {
		dbURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", url.QueryEscape(
			dbUsername), dbPassword, dbHost, dbPort, dbName)
	}

	// 기본 URL 설정
	baseURL := getEnv("SCRAPER_BASE_PATH", "/")
	if baseURL[0] != '/' {
		baseURL = "/" + baseURL
	}

	// Config 구조체 생성
	config := &Config{
		Storage:        storage,
		Media:          media,
		Export:         export,
		DBDriver:       dbDriver,
		DBName:         dbName,
		DBPassword:     dbPassword,
		DBUsername:     dbUsername,
		DBHost:         dbHost,
		DBPort:         dbPort,
		DBURL:          dbURL,
		BaseURL:        baseURL,
		ImageHeightThr: 200,
		ImageWithThr:   200,
	}

	return config
}

// getEnv 함수는 환경 변수를 가져오고, 기본값을 반환합니다.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
