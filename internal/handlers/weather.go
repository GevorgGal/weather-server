package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/GevorgGal/weather-server/internal/data"
	"github.com/GevorgGal/weather-server/internal/models"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")

	if latStr == "" || lonStr == "" {
		http.Error(w, "Missing lat or lon param", http.StatusBadRequest)
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		http.Error(w, "Invalid lat param", http.StatusBadRequest)
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		http.Error(w, "Invalid lon param", http.StatusBadRequest)
		return
	}

	weatherData, err := data.FetchWeatherData(lat, lon)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to fetch weather data: %v", err), http.StatusInternalServerError)
		return
	}

	condition := "Unknown"
	if len(weatherData.Weather) > 0 {
		condition = weatherData.Weather[0].Main
	}
	temperature := "Moderate"
	if weatherData.Main.Temp < 10 {
		temperature = "Cold"
	} else if weatherData.Main.Temp > 25 {
		temperature = "Hot"
	}

	resp := models.WeatherResponse{
		Condition:   condition,
		Temperature: temperature,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("Unable to generate the response: %v", err), http.StatusInternalServerError)
	}
}
