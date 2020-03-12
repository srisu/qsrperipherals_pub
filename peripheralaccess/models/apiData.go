package models

type ApiResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type SocketRequest struct {
	DeviceType    string `json:"deviceType"`
	ActionCode string `json:"actionCode"`
	ReadInterval  string `json:"readInterval"`
}

type SocketResponse struct {
	Value string `json:"value"`
	Status string `json:"status"`
	Unit string `json:"unit"`
	Message string `json:"message"`
}