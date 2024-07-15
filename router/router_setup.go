package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Route struct defines the structure of a route, which includes the route name, HTTP method, pattern, and the handler function to be executed.
type AppRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(*gin.Context)
}

// routes struct holds the Gin engine instance.
type Router struct {
	router *gin.Engine
}

type AppRoutes []AppRoute

// UrlShortner function sets up the URL shortening related routes within a router group.
// It maps each route in the urlShortner slice to its respective HTTP method and handler function
func (r Router) UrlShortenerRoutes(rg *gin.RouterGroup) {
	orderRouteGrouping := rg.Group("/url")
	for _, route := range urlShortner {
		switch route.Method {
		case http.MethodGet:
			orderRouteGrouping.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			orderRouteGrouping.POST(route.Pattern, route.HandlerFunc)
		case http.MethodOptions:
			orderRouteGrouping.OPTIONS(route.Pattern, route.HandlerFunc)
		case http.MethodPut:
			orderRouteGrouping.PUT(route.Pattern, route.HandlerFunc)
		case http.MethodDelete:
			orderRouteGrouping.DELETE(route.Pattern, route.HandlerFunc)
		default:
			orderRouteGrouping.GET(route.Pattern, func(c *gin.Context) {
				c.JSON(200, gin.H{
					"result": "Specify a valid http method with this route.",
				})
			})
		}
	}
}


// ClientRoutes function initializes the router, sets up the routes, and starts the server.
// It uses environment variables to get the API version and port number for the server.
func ClientRoutes() {

	r := Router{
		router: gin.Default(),
	}
	v1 := r.router.Group(os.Getenv("API_VERSION"))
	r.UrlShortenerRoutes(v1)

	if err := r.router.Run(":" + os.Getenv("PORT")); err != nil {
		log.Println("Failed to run server: %v", err)
	}
}
