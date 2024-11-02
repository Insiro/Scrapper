package main

import (
	"Scrapper/cmd/loader"
	"Scrapper/internal/app"
	"Scrapper/pkg/out"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	fmt.Print(os.Args)
	_ = godotenv.Load(".env")
	config := app.InitConfig()
	out.Table(config, "Scrapper Config")
	db := app.InitDB(config)
	if len(os.Args) > 1 {
		switch strings.ToLower(os.Args[1]) {
		case "migrate":
			loader.Migrate(db)
		}
	}
	loader.Web(config, db)

}
