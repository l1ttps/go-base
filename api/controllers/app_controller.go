package controllers

import (
	"go-base/services"

	"github.com/gin-gonic/gin"
)

// AppController is a function that handles the routing for the root URL ("/").
//
// It takes in a *gin.RouterGroup parameter named "r" which represents the router group that the function will be added to.
//
// There are no return types specified for this function.
func AppController() *gin.Engine {
	appService := services.AppService()
	return Controller(ControllerBase{
		basePath: "/hello-world",
		routes: []RouteBase{
			Get("/", appService.HelloWorldGet),

			Post("/", appService.HelloWorldPost),
		},
	})
}
