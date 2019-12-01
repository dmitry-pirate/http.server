package main

import (
	"log"
	"vpn_api/app/server"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	config := server.NewConfig()
	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
