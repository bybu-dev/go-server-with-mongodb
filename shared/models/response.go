package models

type IError struct {
	Field string `json:"field,omitempty"`
	Message string `json:"message"`
}

type IErrors []IError


type IOptions struct {
	Limit int64 `json:"limit"`
	Page int64 `json:"page"`
}

type IErrorResponse struct {
	Status bool `json:"status"`
	Message string `json:"message"`
	Errors IErrors `json:"errors"`
}

type IResponse struct {
	Status bool `json:"status"`
	Data interface{} `json:"data"`
}

func ToSuccessResponse(data interface{}) IResponse {
	return IResponse{
		Status: true,
		Data: data,
	}
}

func ToErrorResponse(errors IErrors) IErrorResponse {
	return IErrorResponse{
		Status: false,
		Message: errors[0].Message,
		Errors: errors,
	}
}