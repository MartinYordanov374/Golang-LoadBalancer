package HelperFunctions

import(
	"encoding/json"
)

type HealthCheckEndpointResponse struct {
	StatusCode int `json:"StatusCode"`
	Message string	`json:"Message"`
}

func ParseJSONResponse(JSONResponse []byte) HealthCheckEndpointResponse {

	var ParsedJSON HealthCheckEndpointResponse
	json.Unmarshal([]byte(JSONResponse), &ParsedJSON)
	return ParsedJSON
}

