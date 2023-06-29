package models

import (
	"encoding/json"
	"errors"
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
		return err
	}
	defer resp.Body.Close()

	var apiResponse IPAPIResponse

	err = json.NewDecoder(resp.Body).Decode(&apiResponse)
	if err != nil {
		return err
	}

	if apiResponse.Status == "fail" {
		if apiResponse.Message == "reserved range" {
			ip.Lat = 0
			ip.Lon = 0
			return nil
		} else {
			return errors.New(apiResponse.Message)
		}
	}

	ip.Lat = apiResponse.Lat
	ip.Lon = apiResponse.Lon

	return nil
}

func (ip *IP) FirstOrCreate(db *gorm.DB) error {
	result := db.First(ip, "ip = ?", ip.IP)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err := ip.CalculateLatLon()
		if err != nil {
			return err
		}
		db.Create(ip)
	} else if result.Error != nil {
		return result.Error
	}
	return nil
}
