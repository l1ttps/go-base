package controllers

import (
	"fmt"
	"go-base/api/exception"
	"go-base/driver"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Engine *gin.Engine

// Controller creates a gin.Engine instance and sets up the routes based on the given basePath and routes.
//
// Parameters:
// - basePath: The base path for all the routes.
// - routes: The list of RouteBase objects representing the routes to be set up.
//
// Returns:
// - *gin.Engine: The gin.Engine instance with the routes set up.
func Controller(basePath string, routes ...RouteBase) *gin.Engine {
	var drive = driver.Instance
	c := drive.Group(basePath)
	methodMap := map[HTTPMethod]func(string, ...gin.HandlerFunc) gin.IRoutes{
		"GET":    c.GET,
		"POST":   c.POST,
		"PUT":    c.PUT,
		"DELETE": c.DELETE,
		"PATCH":  c.PATCH,
	}

	for _, route := range routes {
		fmt.Println(route)
		handlerFunc, exists := methodMap[route.method]
		if !exists {
			fmt.Printf("Invalid HTTP method: %s\n", route.method)
			continue
		}
		// handlerFunc(route.basePath, append(route.middlewares, route.handler)...)
		handlerFunc(route.basePath, append(route.middlewares, PipeResponse(route.handler))...)
	}

	return drive
}

// PipeResponse generates a Gin middleware that handles the response of a given handler function.
//
// The handler function should return an interface{} type. The middleware checks the type of the response
// and performs the appropriate action based on the type.
//
// If the response is of type exception.HttpExceptionResponse, the middleware returns a JSON response
// with the status and message from the exception.
//
// If the response is of type map[string]interface{}, the middleware checks if the map contains a "status"
// field of type int and a "message" field of type string. If both fields exist, the middleware returns
// a JSON response with the status and message from the map.
//
// If the response is of any other type, the middleware returns a JSON response with the response itself.
func PipeResponse(handler func(c *gin.Context) interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := handler(ctx)

		if httpException, ok := response.(exception.HttpExceptionResponse); ok {
			ctx.JSON(httpException.Status, gin.H{
				"status":  httpException.Status,
				"message": httpException.Message,
			})
			return
		}

		if status, ok := response.(map[string]interface{}); ok {
			statusCode, exists := status["status"].(int)
			message, messageExists := status["message"].(string)
			if exists && messageExists {
				ctx.JSON(statusCode, gin.H{
					"status":  statusCode,
					"message": message,
				})
				return
			}
		}

		ctx.JSON(http.StatusOK, response)
	}
}

type RouteBase struct {
	basePath    string
	handler     func(c *gin.Context) any
	method      HTTPMethod
	middlewares []gin.HandlerFunc
}

type MethodHandlerConfigs struct {
	openApi OpenApiConfigs
}

type OpenApiConfigs struct {
	title       string
	description string
	version     string
	tags        []string
	success     string
}

type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	DELETE HTTPMethod = "DELETE"
	PATCH  HTTPMethod = "PATCH"
)

// NewRouteBase creates a new RouteBase struct with the given parameters.
//
// basePath: the base path for the route.
// handler: the handler function for the route.
// method: the HTTP method for the route.
// middlewares: an array of middleware functions for the route.
// Returns: a RouteBase struct.
func NewRouteBase(basePath string, handler func(c *gin.Context) any, method HTTPMethod, middlewares []gin.HandlerFunc) RouteBase {
	return RouteBase{
		basePath:    basePath,
		handler:     handler,
		method:      method,
		middlewares: middlewares,
	}
}

func TestMethod(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, GET, middlewares)
}

// Get returns a new RouteBase with the given base path, handler function, and
// optional middlewares.
//
// Parameters:
//   - basePath: a string representing the base path for the route
//   - handler: a gin.HandlerFunc representing the handler function for the route
//   - middlewares: a variadic list of gin.HandlerFunc representing the optional
//     middlewares for the route
//
// Return:
// - RouteBase: a new instance of RouteBase
func Get(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, GET, middlewares)
}

// Get returns a new RouteBase with the given base path, handler function, and
// optional middlewares.
//
// Parameters:
//   - basePath: a string representing the base path for the route
//   - handler: a gin.HandlerFunc representing the handler function for the route
//   - middlewares: a variadic list of gin.HandlerFunc representing the optional
//     middlewares for the route
//
// Return:
// - RouteBase: a new instance of RouteBase
func Post(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, POST, middlewares)
}

// Put creates a new RouteBase with the specified base path, handler function, and optional middlewares.
//
// Parameters:
// - basePath: the base path for the route.
// - handler: the handler function for the route.
// - middlewares: optional middleware functions to be applied to the route.
//
// Return:
// - RouteBase: a new RouteBase instance.
func Put(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, PUT, middlewares)
}

// Delete returns a new RouteBase with the given base path, handler, DELETE method, and optional middlewares.
//
// Parameters:
// - basePath: The base path for the route.
// - handler: The handler function to be executed for the route.
// - middlewares: Optional middleware functions to be executed before the handler.
//
// Return:
// - RouteBase: The newly created RouteBase.
func Delete(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, DELETE, middlewares)
}

// Patch creates a new RouteBase with the specified base path, handler, and middlewares for the PATCH HTTP method.
//
// Parameters:
// - basePath: The base path for the route.
// - handler: The handler function for the route.
// - middlewares: Optional middlewares to be applied to the route.
//
// Return:
// - RouteBase: The created RouteBase object.
func Patch(basePath string, handler func(c *gin.Context) any, middlewares ...gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, PATCH, middlewares)
}
