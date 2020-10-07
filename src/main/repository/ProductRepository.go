package repository

import (
	. "github.com/hsedjame/products-api/src/main/models"
)

type ProductRepository interface {
	GetAll() Products
	GetById(id int) Product
	GetByName(name string) Product
	Create(product Product) Products
	Update(product Product) Products
	Delete(id int) bool
}
