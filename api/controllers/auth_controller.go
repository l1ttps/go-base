package controllers

import (
	"go-base/api/services"

	"github.com/gin-gonic/gin"
)

// AuthController returns a *gin.Engine instance.
//
// It creates a new instance of the AppService and assigns it to the appService variable.
// Then it returns the result of calling the Controller function with the "/auth" path,
// a Post route with the HelloWorldPost method from the appService variable and a TestGuard
// guard, and a Get route with the HelloWorldGet method from the appService variable.
// The returned *gin.Engine instance is used for handling HTTP requests and responses.
func AuthController() *gin.Engine {
	appService := services.AuthService()
	return Controller("/auth",
		Post("/login", func(c *gin.Context) any {
			return appService.Login()
		},
		),
	)
}
