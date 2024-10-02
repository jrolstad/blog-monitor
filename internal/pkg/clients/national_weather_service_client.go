package clients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type NationalWeatherServiceClient struct {
}

func NewNationalWeatherServiceClient() (*NationalWeatherServiceClient, error) {
	return &NationalWeatherServiceClient{}, nil
}

func (s *NationalWeatherServiceClient) GetProductData(url string) ([]NWSProduct, error) {
	data, err := s.fetchAPIResponse(url)
	if err != nil {
		return []NWSProduct{}, err
	}
	return data.Graph, nil
}

func (s *NationalWeatherServiceClient) GetProductItem(url string) (NWSProduct, error) {
	data, err := s.fetchAPIResponseSingleItem(url)
	if err != nil {
		return NWSProduct{}, err
	}
	return data, nil
}

func (s *NationalWeatherServiceClient) fetchAPIResponse(apiURL string) (NWSApiResponse, error) {
	var response NWSApiResponse

	resp, err := http.Get(apiURL)
	if err != nil {
		return response, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal the JSON response into the Response struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return response, nil
}

func (s *NationalWeatherServiceClient) fetchAPIResponseSingleItem(apiURL string) (NWSProduct, error) {
	var response NWSProduct

	resp, err := http.Get(apiURL)
	if err != nil {
		return response, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return response, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return response, fmt.Errorf("error reading response body: %v", err)
	}

	// Unmarshal the JSON response into the Response struct
	err = json.Unmarshal(body, &response)
	if err != nil {
		return response, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return response, nil
}

type NWSProduct struct {
	Id              string    `json:"id"`
	URL             string    `json:"@id"`
	WmoCollectiveID string    `json:"wmoCollectiveId"`
	IssuingOffice   string    `json:"issuingOffice"`
	IssuanceTime    time.Time `json:"issuanceTime"`
	ProductCode     string    `json:"productCode"`
	ProductName     string    `json:"productName"`
	ProductText     string    `json:"productText"`
}

type NWSApiResponse struct {
	Graph []NWSProduct `json:"@graph"`
}
