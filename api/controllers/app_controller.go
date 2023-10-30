package controllers

import (
	"go-base/api/middlewares/guard"
	"go-base/api/services"

	"github.com/gin-gonic/gin"
)

var (
	UseGuard  = guard.UseGuard
	TestGuard = guard.TestGuard
)

// AppController is a function that handles the routing for the root URL ("/").
//
// It takes in a *gin.RouterGroup parameter named "r" which represents the router group that the function will be added to.
//
// There are no return types specified for this function.
func AppController() *gin.Engine {

	appService := services.AppService()
	return Controller("/",
		Get("/hello-world",
			func(c *gin.Context) interface{} {
				return appService.HelloWorldGet()
			},
			UseGuard(TestGuard),
		),

		Get("/forbidden",
			func(c *gin.Context) interface{} {
				return appService.Forbidden()
			},
		),
	)
}
