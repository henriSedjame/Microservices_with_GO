package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

type ProductController struct {
	handler ProductHandler
}

func NewProductController(handler ProductHandler) *ProductController {
	return &ProductController{handler: handler}
}

func (pCtrl ProductController) Path() string {
	return "/products"
}

func (pCtrl ProductController) AddRoutes(router *mux.Router) {

	router.
		HandleFunc("/", pCtrl.handler.GetProducts).
		Methods(http.MethodGet)

	router.
		HandleFunc("/", pCtrl.handler.CreateProduct).
		Methods(http.MethodPost)

	router.
		HandleFunc("/{id:[0-9]+}", pCtrl.handler.GetProduct).
		Methods(http.MethodGet)

	router.
		HandleFunc("/{name}", pCtrl.handler.GetProduct).
		Methods(http.MethodGet)

	router.
		HandleFunc("/", pCtrl.handler.UpdateProduct).
		Methods(http.MethodPut)

	router.
		HandleFunc("/", pCtrl.handler.DeleteProduct).
		Methods(http.MethodDelete)
}

