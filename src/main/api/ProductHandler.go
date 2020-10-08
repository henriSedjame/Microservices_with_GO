package api

import (
	"github.com/gorilla/mux"
	"github.com/hsedjame/products-api/src/main/core"
	. "github.com/hsedjame/products-api/src/main/models"
	. "github.com/hsedjame/products-api/src/main/repository"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	repoditory ProductRepository
}

func NewProductHandler(productRepository ProductRepository) *ProductHandler {
	return &ProductHandler{repoditory: productRepository}
}

// swagger:operation GET /products products listOfProducts
//
// Retrieve list of all available products
//
// ---
//	consumes:
//	- application/json
//
//	produces:
//	- application/json
//
//
// responses:
//   '200':
//     description: product found
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Product"
//   '400':
//     description: bad request
//     schema:
//       "$ref": "#/definitions/GenericError"
func (handler *ProductHandler) GetProducts(wr http.ResponseWriter, _ *http.Request) {

	products := handler.repoditory.GetAll()
	if err := products.ToJson(wr); err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		_ = core.ToJson(GenericError{Message: err.Error()}, wr)
		return
	}
	return
}

func (handler *ProductHandler) GetProductById(wr http.ResponseWriter, rq *http.Request) {

	pathParams := mux.Vars(rq)

	if idString := pathParams["id"]; idString != "" {
		if id, err := strconv.Atoi(idString); err != nil {
			http.Error(wr, "", http.StatusBadRequest)
			return
		} else {
			product := handler.repoditory.GetById(id)
			if err := core.ToJson(product, wr); err != nil {
				wr.WriteHeader(http.StatusBadRequest)
				_ = core.ToJson(GenericError{Message: err.Error()}, wr)
				return
			}
			return
		}

	}

	wr.WriteHeader(http.StatusBadRequest)
	_ = core.ToJson(GenericError{Message: "Parameter 'id' required"}, wr)
	return
}

// swagger:operation GET /products/{name} productByName getProductByName
//
// Retrieve a specific product by its name
//
// ---
//	consumes:
//	- application/json
//
//	produces:
//	- application/json
//
// parameters:
// - name: name
//   in: path
//   description: name of product
//   required: true
//   type: string
//
//
// responses:
//   '200':
//     description: product found
//     schema:
//       "$ref": "#/definitions/Product"
//
//   '400':
//     description: bad request
//     schema:
//       "$ref": "#/definitions/GenericError"
func (handler *ProductHandler) GetProductByName(wr http.ResponseWriter, rq *http.Request) {

	pathParams := mux.Vars(rq)

	if name := pathParams["name"]; name != "" {
		product := handler.repoditory.GetByName(name)
		if err := core.ToJson(product, wr); err != nil {
			wr.WriteHeader(http.StatusBadRequest)
			_ = core.ToJson(GenericError{Message: err.Error()}, wr)
			return
		}
		return
	}

	wr.WriteHeader(http.StatusBadRequest)
	_ = core.ToJson(GenericError{Message: "Parameter 'name' required "}, wr)
	return
}

// swagger:operation POST /products createProduct createProduct
//
// Create a new product
//
// ---
//	consumes:
//	- application/json
//
//	produces:
//	- application/json
//
// parameters:
// - name: body
//   in: body
//   description: body of request
//   required: true
//   schema:
//          $ref: '#/definitions/Product'
//
//
// responses:
//   '200':
//     description: product found
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Product"
//   '400':
//     description: bad request
//     schema:
//       "$ref": "#/definitions/GenericError"
//   '406':
//     description: not acceptable error
//     schema:
//       "$ref": "#/definitions/ValidationError"
//   '500':
//     description: internal server error
//     schema:
//       "$ref": "#/definitions/GenericError"
func (handler *ProductHandler) CreateProduct(wr http.ResponseWriter, rq *http.Request) {

	prod := rq.Context().Value(ProductKey{}).(Product)

	products := handler.repoditory.Create(prod)

	if err := products.ToJson(wr); err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		_ = core.ToJson(GenericError{Message: err.Error()}, wr)
		return
	}

	return
}

// swagger:operation PUT /products updateProduct updateProduct
//
// Update a product
//
// ---
//	consumes:
//	- application/json
//
//	produces:
//	- application/json
//
// parameters:
// - name: body
//   in: body
//   description: body of request
//   required: true
//   schema:
//          $ref: '#/definitions/Product'
//
//
// responses:
//   '200':
//     description: product updated successfully
//     schema:
//       type: array
//       items:
//         "$ref": "#/definitions/Product"
//   '400':
//     description: bad request
//     schema:
//       "$ref": "#/definitions/GenericError"
//   '406':
//     description: request not acceptable error
//     schema:
//       "$ref": "#/definitions/ValidationError"
//   '500':
//     description: internal server error
//     schema:
//       "$ref": "#/definitions/GenericError"
func (handler *ProductHandler) UpdateProduct(wr http.ResponseWriter, rq *http.Request) {

	prod := rq.Context().Value(ProductKey{}).(Product)

	products := handler.repoditory.Update(prod)

	if err := products.ToJson(wr); err != nil {
		wr.WriteHeader(http.StatusBadRequest)
		_ = core.ToJson(GenericError{Message: err.Error()}, wr)
		return
	}

	return
}

// swagger:operation DELETE /products/{id} deleteProduct deleteProduct
//
// Delete a product
//
// ---
//	consumes:
//	- application/json
//
//	produces:
//	- application/json
//
// parameters:
// - name: id
//   in: path
//   description: id of product to delete
//   required: true
//   type: integer
//
// responses:
//   '200':
//     description: product updated successfully
//     schema:
//       type: boolean
//   '500':
//     description: internal server error
//     schema:
//       "$ref": "#/definitions/GenericError"
func (handler *ProductHandler) DeleteProduct(wr http.ResponseWriter, rq *http.Request) {

	pathParams := mux.Vars(rq)

	if idString := pathParams["id"]; idString != "" {

		if id, err := strconv.Atoi(idString); err != nil {

		} else {
			if handler.repoditory.Delete(id) {
				wr.WriteHeader(http.StatusOK)
				return
			} else {
				wr.WriteHeader(http.StatusInternalServerError)
				_ = core.ToJson(GenericError{Message: "Une erreur inatendue s'est produite"}, wr)
				return
			}
		}
	}
	return
}
