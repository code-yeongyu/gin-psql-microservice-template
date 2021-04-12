package routes

import "github.com/gin-gonic/gin"

// RouteInfo is a struct for a route
type RouteInfo struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}
