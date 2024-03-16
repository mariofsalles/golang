package router

import (
	"webapp/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a new router configured
func Generate() *mux.Router {
	return routes.Configure(mux.NewRouter())
}
