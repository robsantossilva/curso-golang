package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func BuildRouter() *mux.Router {
	r := mux.NewRouter()
	return routes.Setup(r)
}
