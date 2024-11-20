package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Struct to hold the weather data
type WeatherResponse struct {
	Main struct {
		Temp     float64 `json:"temp"`
		Humidity int     `json:"humidity"`
	} `json:"main"`
	Name string `json:"name"`
}

func getWeather(city string, apiKey string) (*WeatherResponse, error) {
	// Construct the API URL
	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parse JSON response into WeatherResponse struct
	var weather WeatherResponse
	if err := json.Unmarshal(body, &weather); err != nil {
		return nil, err
	}

	return &weather, nil
}

func main() {
	// Replace with your OpenWeatherMap API Key
	apiKey := "YOUR_API_KEY_HERE"
	city := "London" // City for which you want the weather

	// Get weather data
	weather, err := getWeather(city, apiKey)
	if err != nil {
		log.Fatalf("Error getting weather data: %v", err)
	}

	// Output the weather data
	fmt.Printf("Weather in %s:\n", weather.Name)
	fmt.Printf("Temperature: %.2f°C\n", weather.Main.Temp)
	fmt.Printf("Humidity: %d%%\n", weather.Main.Humidity)
}
