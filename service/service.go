package service

import (
	"encoding/json"
	result "github.com/heaptracetechnology/amplitude/result"
	amplitude "github.com/msingleton/amplitude-go"
	"net/http"
	"os"
)

//EventArguments struct
type EventArguments struct {
	AppVersion         string                 `json:"appVersion,omitempty"`
	Carrier            string                 `json:"carrier,omitempty"`
	City               string                 `json:"city,omitempty"`
	Country            string                 `json:"country,omitempty"`
	DeviceBrand        string                 `json:"deviceBrand,omitempty"`
	DeviceID           string                 `json:"deviceId,omitempty"`
	DeviceManufacturer string                 `json:"deviceManufacturer,omitempty"`
	DeviceModel        string                 `json:"deviceModel,omitempty"`
	EventProperties    map[string]interface{} `json:"eventProperties,omitempty"`
	EventType          string                 `json:"eventType,omitempty"`
	Groups             map[string]interface{} `json:"groups,omitempty"`
	IP                 string                 `json:"ip,omitempty"`
	Language           string                 `json:"language,omitempty"`
	LocationLat        string                 `json:"locationLat,omitempty"`
	LocationLng        string                 `json:"locationLng,omitempty"`
	OsName             string                 `json:"osName,omitempty"`
	OsVersion          string                 `json:"osVersion,omitempty"`
	Platform           string                 `json:"platform,omitempty"`
	Price              float64                `json:"price,omitempty"`
	ProductID          string                 `json:"productId,omitempty"`
	Quantity           int                    `json:"quantity,omitempty"`
	Region             string                 `json:"region,omitempty"`
	Revenue            float64                `json:"revenue,omitempty"`
	RevenueType        string                 `json:"revenueType,omitempty"`
	UserID             string                 `json:"userId,omitempty"`
	UserProperties     map[string]interface{} `json:"userProperties,omitempty"`
}

//Message struct
type Message struct {
	Success    string `json:"success"`
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

//UploadEvent Amplitude
func UploadEvent(responseWriter http.ResponseWriter, request *http.Request) {

	apiKey := os.Getenv("API_KEY")

	decoder := json.NewDecoder(request.Body)
	var param EventArguments
	decodeErr := decoder.Decode(&param)
	if decodeErr != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	client := amplitude.New(apiKey)
	err := client.Event(amplitude.Event{
		AppVersion:         param.AppVersion,
		Carrier:            param.Carrier,
		City:               param.City,
		Country:            param.Country,
		DeviceBrand:        param.DeviceBrand,
		DeviceId:           param.DeviceID,
		DeviceManufacturer: param.DeviceManufacturer,
		DeviceModel:        param.DeviceModel,
		EventProperties:    param.EventProperties,
		EventType:          param.EventType,
		Groups:             param.Groups,
		Ip:                 param.IP,
		Language:           param.Language,
		LocationLat:        param.LocationLat,
		LocationLng:        param.LocationLng,
		OsName:             param.OsName,
		OsVersion:          param.OsVersion,
		Platform:           param.Platform,
		Price:              param.Price,
		ProductId:          param.ProductID,
		Quantity:           param.Quantity,
		Region:             param.Region,
		Revenue:            param.Revenue,
		RevenueType:        param.RevenueType,
		UserId:             param.UserID,
		UserProperties:     param.UserProperties,
	})
	if err != nil {
		result.WriteErrorResponseString(responseWriter, decodeErr.Error())
		return
	}

	message := Message{"true", "Event uploaded successfully", http.StatusOK}
	bytes, _ := json.Marshal(message)
	result.WriteJsonResponse(responseWriter, bytes, http.StatusOK)
}
