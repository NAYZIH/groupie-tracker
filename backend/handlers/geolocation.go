package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetCoordinates(location string) (float64, float64, error) {
	location = strings.ReplaceAll(location, " ", "+")
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/search?format=json&q=%s", location)

	resp, err := http.Get(url)
	if err != nil {
		return 0, 0, err
	}
	defer resp.Body.Close()

	var results []map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&results); err != nil {
		return 0, 0, err
	}

	if len(results) == 0 {
		return 0, 0, fmt.Errorf("no results found for location: %s", location)
	}

	lon, ok := results[0]["lon"].(string)
	if !ok {
		return 0, 0, fmt.Errorf("unable to extract longitude")
	}
	lat, ok := results[0]["lat"].(string)
	if !ok {
		return 0, 0, fmt.Errorf("unable to extract latitude")
	}

	var latitude, longitude float64
	longitude, err = strconv.ParseFloat(lon, 64)
	if err != nil {
		return 0, 0, err
	}

	latitude, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		return 0, 0, err
	}

	return longitude, latitude, nil
}
