package services

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

func (AppService *AppServiceType) HelloWorldGet() HelloWorldResponse {
	return HelloWorldResponse{
		Message: "[GET] Hello World",
	}
}
