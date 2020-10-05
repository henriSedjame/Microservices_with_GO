package api

import (
	"github.com/gorilla/mux"
	"github.com/hsedjame/products-api/src/models"
	"github.com/hsedjame/products-api/src/repository"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	repoditory repository.ProductRepository
}

func NewProductHandler(productRepository repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repoditory: productRepository}
}

func (handler *ProductHandler) GetProducts(wr http.ResponseWriter, _ *http.Request)  {
	products := handler.repoditory.GetAll()
	if err:=products.ToJson(wr); err !=nil {
		http.Error(wr, "", http.StatusBadRequest)
		return
	}
	return
}

func (handler *ProductHandler) GetProduct(wr http.ResponseWriter, rq *http.Request)  {

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

	} else if name := pathParams["name"]; name != "" {
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

func (handler *ProductHandler) CreateProduct(wr http.ResponseWriter, rq *http.Request)  {

	var prod models.Product

	if err := prod.FromJson(rq.Body); err != nil {
		http.Error(wr, "", http.StatusBadRequest)
		return
	}
	products := handler.repoditory.Create(prod)

	if err:=products.ToJson(wr); err !=nil {
		http.Error(wr, "", http.StatusBadRequest)
		return
	}

	return
}

func (handler *ProductHandler) UpdateProduct(wr http.ResponseWriter, rq *http.Request)  {

	var prod models.Product

	if err := prod.FromJson(rq.Body); err != nil {
		http.Error(wr, "", http.StatusBadRequest)
		return
	}
	products := handler.repoditory.Update(prod)

	if err:=products.ToJson(wr); err !=nil {
		http.Error(wr, "", http.StatusBadRequest)
		return
	}

	return
}

func (handler *ProductHandler) DeleteProduct(wr http.ResponseWriter, rq *http.Request)  {

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
