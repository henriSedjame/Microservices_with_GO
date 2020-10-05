package api

import "github.com/gorilla/mux"

type RestController interface {
	Path() string
	AddRoutes(router *mux.Router)
}
