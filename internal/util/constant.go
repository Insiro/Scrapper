package util

import "Scrapper/pkg/utils"

const USER_AGENT = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36"

var IMAGE_EXTENSIONS = utils.NewSet[string](".jpg", ".jpeg", ".png", ".gif", ".webp", ".svg")
