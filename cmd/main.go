package main

import (
	"Scrapper/cmd/loader"
	"Scrapper/internal/app"
	"Scrapper/pkg/out"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	var command = ""
	args := os.Args
	argCount := len(args)

	if argCount > 1 {
		command = strings.ToLower(args[1])
	}
	_ = godotenv.Load(".env")

	config := app.InitConfig()
	out.Table(config, "Scrapper Config")

	db := app.InitDB(config)
	switch command {
	case "migrate":
		loader.Migrate(db)
		os.Exit(1)
	case "scrap":
		loader.Scrap(args[2], db, config)
		os.Exit(1)
	}
	loader.Web(config, db)
}
