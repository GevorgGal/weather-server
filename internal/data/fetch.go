package data

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GevorgGal/weather-server/config"
	"github.com/GevorgGal/weather-server/internal/models"
)

func FetchWeatherData(lat, lon float64) (*models.WeatherData, error) {
	apiKey := config.GetAPIKey()
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%.2f&lon=%.2f&appid=%s&units=metric", lat, lon, apiKey)

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %v", err)
	}
	defer res.Body.Close()

	var data models.WeatherData
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("unable to parse weather data: %v", err)
	}
	return &data, nil
}
