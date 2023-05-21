package models

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Species struct {
	ID            uint     `json:"id" gorm:"primaryKey"`
	CommonName    string   `json:"common_name" gorm:"not null"`
	Scientific    string   `json:"scientific" gorm:"not null"`
	SpeciesAdvice []Advice `json:"species_advice" gorm:"foreignKey:SpeciesID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (s *Species) GetData() error {
	// Get the Trefle API key from environment variables
	api_key := os.Getenv("TREFLE_API_KEY")

	// Construct the URL for the Trefle API request
	url := fmt.Sprintf("https://trefle.io/api/v1/plants/search?token=%s&q=%s", api_key, s.CommonName)

	// Make a request to the Trefle API
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	type TrefleResponse struct {
		Data []struct {
			ScientificName string `json:"scientific_name"`
			CommonName     string `json:"common_name"`
		} `json:"data"`
	}

	// Decode the response
	var trefleResp TrefleResponse
	err = json.NewDecoder(resp.Body).Decode(&trefleResp)
	if err != nil {
		return err
	}

	// Update the species instance with the fetched data
	// This assumes that the first result in the Trefle API response is the correct species
	if len(trefleResp.Data) > 0 {
		s.Scientific = trefleResp.Data[0].ScientificName
		s.CommonName = trefleResp.Data[0].CommonName
	}

	return nil
}
