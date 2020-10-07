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
