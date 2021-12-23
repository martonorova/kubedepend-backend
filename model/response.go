package model

type APIResponseStatus string

const (
	SuccessStatus APIResponseStatus = "success"
	ErrorStatus   APIResponseStatus = "error"
)

type APIResponse struct {
	Status  APIResponseStatus `json:"status"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data"`
}

// TODO should use pointers instead?
func APISuccess(data interface{}) APIResponse {
	return APIResponse{Status: SuccessStatus, Data: data}
}

func APIError() APIResponse {
	return APIResponse{Status: ErrorStatus}
}

func APIErrorWithData(data interface{}) APIResponse {
	return APIResponse{Status: ErrorStatus, Data: data}
}
