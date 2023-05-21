package models

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Location struct {
	gorm.Model
	Name      string  `json:"name" gorm:"type:varchar(255);not null"`
	Address   string  `json:"address" gorm:"type:varchar(255);not null"`
	City      string  `json:"city" gorm:"type:varchar(255);not null"`
	ZipCode   string  `json:"zip_code" gorm:"type:varchar(255);not null"`
	Country   string  `json:"country" gorm:"type:varchar(255);not null"`
	Latitude  float64 `json:"latitude" gorm:"not null"`
	Longitude float64 `json:"longitude" gorm:"not null"`
	UserID    uint    `json:"user_id" gorm:"not null"`
}

func (l Location) CalculateLatLon() (float64, float64, error) {
	address := l.Address + " " + l.City + " " + l.ZipCode + " " + l.Country

	baseURL := "https://nominatim.openstreetmap.org/search"
	u, err := url.Parse(baseURL)
	if err != nil {
		return 0, 0, err
	}

	params := url.Values{}
	params.Add("q", address)
	params.Add("format", "json")
	u.RawQuery = params.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return 0, 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var responseJson []map[string]interface{}
	err = json.Unmarshal(body, &responseJson)
	if err != nil {
		return 0, 0, err
	}

	if len(responseJson) > 0 {
		firstItem := responseJson[0]

		latString := firstItem["lat"]
		lonString := firstItem["lon"]

		lat, err := strconv.ParseFloat(latString.(string), 64)
		lon, err := strconv.ParseFloat(lonString.(string), 64)

		l.Latitude = lat
		l.Longitude = lon

		if err != nil {
			return 0, 0, err
		}

		return lat, lon, nil
	}

	fmt.Println("No data available")
	return 0, 0, err
}

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

// Create a new location
func (r *LocationRepository) Create(location *Location) error {
	return r.db.Create(location).Error
}

// Get a location by ID
func (r *LocationRepository) GetByID(id uint) (*Location, error) {
	var location Location
	err := r.db.First(&location, id).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

// Update a location
func (r *LocationRepository) Update(location *Location) error {
	return r.db.Save(location).Error
}

// Delete a location
func (r *LocationRepository) Delete(location *Location) error {
	return r.db.Delete(location).Error
}

// Get location by user ID
func (r *LocationRepository) GetByUserID(userID uint) (*Location, error) {
	var location Location
	err := r.db.Where("user_id = ?", userID).First(&location).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}
