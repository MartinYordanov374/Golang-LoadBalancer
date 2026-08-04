package HelperFunctions

import(
	"encoding/json"
	"golang-loadbalancer/Structs"
)

func ParseJSONResponse(JSONResponse []byte) Structs.HealthCheckEndpointResponse {
	var ParsedJSON Structs.HealthCheckEndpointResponse
	json.Unmarshal([]byte(JSONResponse), &ParsedJSON)
	return ParsedJSON
}
