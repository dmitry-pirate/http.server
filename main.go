package main

import (
	"github.com/dmitry-pirate/http.server/app"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No dot env file")
	}
}

func main() {
	api := app.New()
	if err := api.Start(); err != nil {
		panic(err)
	}
}
