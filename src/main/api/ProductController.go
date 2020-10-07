// Package api Product API.
//
// the purpose of this application is to provide an application
// that is using plain go code to define an API
//
// This should demonstrate all the possible comment annotations
// that are available to turn go code into a fully compliant swagger 2.0 spec
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http
//     Host: localhost
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//     - api_key:
//
// swagger:meta
package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/hsedjame/products-api/src/main/models"
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
		HandleFunc("", pCtrl.handler.GetProducts).
		Methods(http.MethodGet)

	router.
		HandleFunc("", pCtrl.handler.CreateProduct).
		Methods(http.MethodPost)

	router.
		HandleFunc("/{id:[0-9]+}", pCtrl.handler.GetProductById).
		Methods(http.MethodGet)

	router.
		HandleFunc("/{name}", pCtrl.handler.GetProductByName).
		Methods(http.MethodGet)

	router.
		HandleFunc("", pCtrl.handler.UpdateProduct).
		Methods(http.MethodPut)

	router.
		HandleFunc("", pCtrl.handler.DeleteProduct).
		Methods(http.MethodDelete)
}

func (pCtrl ProductController) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, rq *http.Request) {

		if rq.Method == http.MethodPost {
			var prod models.Product
			if err := prod.FromJson(rq.Body); err != nil {
				http.Error(wr, fmt.Sprintf("Invalid request %s", err), http.StatusBadRequest)
			} else if err := prod.Validate(); err != nil {
				http.Error(wr, fmt.Sprintf("Invalid request %s", err), http.StatusBadRequest)
			}
		}

		next.ServeHTTP(wr, rq)
	})
}
