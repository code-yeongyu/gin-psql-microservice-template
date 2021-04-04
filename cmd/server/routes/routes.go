package routes

import (
	"github.com/gin-gonic/gin"
)

// RouteInfo is a struct for a route
type RouteInfo struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// GetRoutes is a function which returns the information of routes
func GetRoutes() []RouteInfo {
	//controller := controllers.Controller{}
	routeInfo := []RouteInfo{}
	return routeInfo
}
