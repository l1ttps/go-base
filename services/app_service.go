package services

import (
	"go-base/internal/exception"
	"net/http"
)

type AppServiceType struct {
}

func AppService() *AppServiceType {
	return &AppServiceType{}
}

type HelloWorldResponse struct {
	Message string `json:"message"`
}

func (AppService *AppServiceType) HelloWorldPost() HelloWorldResponse {
	return HelloWorldResponse{
		Message: "[POST] Hello World",
	}
}

func (AppService *AppServiceType) HelloWorldGet() interface{} {
	return HelloWorldResponse{
		Message: "[GET] Hello World",
	}
}

func (AppService *AppServiceType) Forbidden() interface{} {
	return exception.HttpException(http.StatusForbidden, "Forbidden")
}
