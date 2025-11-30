package models

type WeatherEvent struct {
	Timestamp       string  `json:"timestamp"`
	Temperature     float64 `json:"temperature"`
	Humidity        float64 `json:"humidity"`
	WindSpeed       float64 `json:"wind_speed"`
	RainProbability float64 `json:"rain_probability"`
	Lat             string  `json:"lat"`
	Lon             string  `json:"lon"`
}
