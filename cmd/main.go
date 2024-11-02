package main

import (
	"Scrapper/cmd/loader"
	"Scrapper/internal/app"
	"Scrapper/internal/scrapper"
	"Scrapper/pkg/out"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	var command = ""
	if len(os.Args) > 1 {
		command = strings.ToLower(os.Args[1])
	}
	_ = godotenv.Load(".env")
	config := app.InitConfig()

	out.Table(config, "Scrapper Config")
	db := app.InitDB(config)
	switch command {
	case "migrate":
		loader.Migrate(db)
		return
	case "scrap":
		src, err := scrapper.Factory("https://x.com/CherryPie_85/status/1772978070942568837?t=a0L0kpQbeEQusZfL9vMDGQ&s=19", nil, nil, config)
		if err != nil {
			panic(err.Error())
		}
		scrap, err := src.Scrap(&src.Args)
		if err != nil {
			return
		}
		fmt.Print(scrap)
		return
	}
	loader.Web(config, db)

}
