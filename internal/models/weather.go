package models

type WeatherData struct {
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

type WeatherResponse struct {
	Condition   string `json:"condition"`
	Temperature string `json:"temperature"`
}
