package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type IP struct {
	gorm.Model
	IP  string  `json:"ip"`
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type IPAPIResponse struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Status  string  `json:"status"`
	Message string  `json:"message"`
}

func (ip *IP) CalculateLatLon() error {
	resp, err := http.Get("http://ip-api.com/json/" + ip.IP)
	if err != nil {
		return fmt.Errorf("failed to make HTTP request to IP-API: %w", err)
	}
	defer resp.Body.Close()

	var apiResponse IPAPIResponse

	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return fmt.Errorf("failed to decode IP-API response: %w", err)
	}

	if apiResponse.Status == "fail" {
		if apiResponse.Message == "reserved range" || apiResponse.Message == "private range" {
			ip.Lat = 0
			ip.Lon = 0
			return nil
		} else {
			return fmt.Errorf("IP-API returned fail status: %s", apiResponse.Message)
		}
	}

	ip.Lat = apiResponse.Lat
	ip.Lon = apiResponse.Lon

	return nil
}

func (ip *IP) FirstOrCreate(db *gorm.DB) (*IP, error) {
	result := db.First(ip, "ip = ?", ip.IP)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err := ip.CalculateLatLon()
		if err != nil {
			return nil, fmt.Errorf("failed to calculate Lat and Lon for IP %s: %w", ip.IP, err)
		}
		result = db.Create(ip)
		if result.Error != nil {
			return nil, fmt.Errorf("failed to create IP record in database: %w", result.Error)
		}
	} else if result.Error != nil {
		return nil, fmt.Errorf("database error when trying to find IP record: %w", result.Error)
	}
	return ip, nil
}
