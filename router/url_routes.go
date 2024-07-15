package router

import (
	"net/http"
	"urlshortner/constant"
	"urlshortner/controller"
)

// urlShortner defines the routes for the URL shortening service.
// Each route is represented by a Route struct which includes the
// route name,
// HTTP method,
// pattern, and the handler function to be executed.
var urlShortner = AppRoutes{
	// Route for shortening a URL.
	// Method: POST
	// Pattern: /short
	// Handler: controller.ShortTheUrl
	AppRoute{"Url Shortner Service", http.MethodPost, constant.UrlShortnerPath, controller.ShortTheUrl},

	// Route for redirecting to the original URL based on the shortened code.
	// Method: GET
	// Pattern: /shubhurl/:code
	// Handler: controller.RedirectURL
	AppRoute{"Redirect to url", http.MethodGet, constant.RedirectUrlPath, controller.RedirectURL},
}
