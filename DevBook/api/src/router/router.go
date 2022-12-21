package router

import "github.com/gorilla/mux"

func BuildRouter() *mux.Router {
	return mux.NewRouter()
}
