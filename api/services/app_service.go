package services

import exception "go-base/api/exception"

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
	return exception.ForbiddenException()
}
