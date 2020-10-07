package api

import (
	"github.com/gorilla/mux"
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
func (handler *ProductHandler) GetProducts(wr http.ResponseWriter, _ *http.Request) {
	products := handler.repoditory.GetAll()
	if err := products.ToJson(wr); err != nil {
		http.Error(wr, "", http.StatusBadRequest)
		return
	}
	return
}

// swagger:operation GET /products/{id}  productById getProductById
//
// Retrieve a specific product by its id
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
//   description: id of product
//   required: true
//   type: integer
//
//
// responses:
//   '200':
//     description: product found
//     schema:
//       "$ref": "#/definitions/Product"
//
func (handler *ProductHandler) GetProductById(wr http.ResponseWriter, rq *http.Request) {

	pathParams := mux.Vars(rq)

	if idString := pathParams["id"]; idString != "" {
		if id, err := strconv.Atoi(idString); err != nil {
			http.Error(wr, "", http.StatusBadRequest)
			return
		} else {
			product := handler.repoditory.GetById(id)
			if err := product.ToJson(wr); err != nil {
				http.Error(wr, "", http.StatusBadRequest)
				return
			}
			return
		}

	}

	http.Error(wr, "", http.StatusBadRequest)
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
func (handler *ProductHandler) GetProductByName(wr http.ResponseWriter, rq *http.Request) {

	pathParams := mux.Vars(rq)

	if name := pathParams["name"]; name != "" {
		product := handler.repoditory.GetByName(name)
		if err := product.ToJson(wr); err != nil {
			http.Error(wr, "", http.StatusBadRequest)
			return
		}
		return
	}

	http.Error(wr, "", http.StatusBadRequest)
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
func (handler *ProductHandler) CreateProduct(wr http.ResponseWriter, rq *http.Request) {

	var prod Product

	if err := prod.FromJson(rq.Body); err != nil {
		http.Error(wr, "", http.StatusBadRequest)
		return
	}
	products := handler.repoditory.Create(prod)

	if err := products.ToJson(wr); err != nil {
		http.Error(wr, "", http.StatusBadRequest)
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
func (handler *ProductHandler) UpdateProduct(wr http.ResponseWriter, rq *http.Request) {

	var prod Product

	if err := prod.FromJson(rq.Body); err != nil {
		http.Error(wr, "", http.StatusBadRequest)
		return
	}
	products := handler.repoditory.Update(prod)

	if err := products.ToJson(wr); err != nil {
		http.Error(wr, "", http.StatusBadRequest)
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

func (handler *ProductHandler) DeleteProduct(wr http.ResponseWriter, rq *http.Request) {

	pathParams := mux.Vars(rq)

	if idString := pathParams["id"]; idString != "" {

		if id, err := strconv.Atoi(idString); err != nil {

		} else {
			if handler.repoditory.Delete(id) {
				wr.WriteHeader(http.StatusOK)
				return
			} else {
				http.Error(wr, "", http.StatusInternalServerError)
				return
			}
		}
	}
	return
}
