package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a new router configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
