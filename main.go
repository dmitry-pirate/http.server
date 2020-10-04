package main

import (
	"context"
	"github.com/basketforcode/http.server/app"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No dot env file")
	}
}

func main() {
	api := app.New()

	go func() {
		if err := api.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := api.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
