package controllers

import (
	"go-base/driver"

	"github.com/gin-gonic/gin"
)

type ControllerBase struct {
	basePath string
	routes   []RouteBase
}

func Controller(params ControllerBase) *gin.Engine {
	var drive = driver.Instance
	c := drive.Group(params.basePath)
	methodMap := map[HTTPMethod]func(string, ...gin.HandlerFunc) gin.IRoutes{
		"GET":    c.GET,
		"POST":   c.POST,
		"PUT":    c.PUT,
		"DELETE": c.DELETE,
		"PATCH":  c.PATCH,
	}

	for _, route := range params.routes {
		if handlerFunc, exists := methodMap[route.method]; exists {
			handlerFunc(route.basePath, route.handler)
		}
	}

	return drive
}

type RouteBase struct {
	basePath string
	handler  gin.HandlerFunc
	method   HTTPMethod
}

type MethodBase struct {
	basePath string
	handler  gin.HandlerFunc
	configs  MethodHandlerConfigs
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

func NewRouteBase(basePath string, handler gin.HandlerFunc, method HTTPMethod) RouteBase {
	return RouteBase{
		basePath: basePath,
		handler:  handler,
		method:   method,
	}
}

func Get(basePath string, handler gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, GET)
}

func Post(basePath string, handler gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, POST)
}

func Put(basePath string, handler gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, PUT)
}

func Delete(basePath string, handler gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, DELETE)
}

func Patch(basePath string, handler gin.HandlerFunc) RouteBase {
	return NewRouteBase(basePath, handler, PATCH)
}
