package routes

import (
	"net/http"

	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Route represents all routes from the application
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

var allRoutes = [][]Route{{homePageRoute}, {logoutRoute}, usersRoutes, loginRoutes, postsRoutes}

// Configure put all routes in the router
func Configure(router *mux.Router) *mux.Router {
	var routes []Route
	for _, routeSlice := range allRoutes {
		routes = append(routes, routeSlice...)
	}

	for _, route := range routes {
		if route.AuthRequired {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(middlewares.Authentication(route.Function)),
			).Methods(route.Method)
		}
		if !route.AuthRequired {
			router.HandleFunc(
				route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
