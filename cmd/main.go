package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GevorgGal/weather-server/internal/handlers"
)

func main() {
	http.HandleFunc("/weather", handlers.WeatherHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}