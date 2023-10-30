package controllers

import (
	guard "go-base/api/middlewares/guards"
	"go-base/services"
)

// AppController is a function that handles the routing for the root URL ("/").
//
// It takes in a *gin.RouterGroup parameter named "r" which represents the router group that the function will be added to.
//
// There are no return types specified for this function.t
func AppController() Engine {
	appService := services.AppService()
	return Controller("/hello-world",
		Get("/", func() any {
			return appService.HelloWorldGet()
		},
			// Implement Guard for Method Get Hello World
			guard.UseGuard(guard.TestGuard),
		),

		Post("/", func() any {
			println("hello")
			return appService.HelloWorldPost()
		},
		),
	)
}
