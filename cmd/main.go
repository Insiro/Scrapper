package main

import (
	"Scrapper/cmd/loader"
	"Scrapper/internal/util"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	fmt.Print(os.Args)
	_ = godotenv.Load(".env")
	config := util.InitConfig()
	db := util.InitDB(config)
	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[0]) {
		case "migrate":
			loader.Migrate(db)
		}
	}
	loader.Web(config)

}
