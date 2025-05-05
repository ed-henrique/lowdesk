package main

import (
	"log/slog"
	"os"

	"github.com/ed-henrique/lowdesk/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("could not load .env file", slog.String("err", err.Error()))
		os.Exit(1)
	}

	s := server.New(server.Config{
		Addr:  ":8080",
		IsDev: os.Getenv("GO_ENV") != "production",
	})

	s.AddRoutes()
	s.Run()
}
