package apiserver

type ApiResponseData struct {
	Items           []string `json:"items"`
	AreAllEncrypted bool     `json:"areAllEncrypted"`
}

type ApiResponse struct {
	Code    int32           `json:"code"`
	Message string          `json:"message"`
	Data    ApiResponseData `json:"data"`
}
