package impl

import (
	. "github.com/hsedjame/products-api/src/main/models"
	"log"
	"sort"
	"time"
)

type SimpleProductRepository struct {
	logger   *log.Logger
	products Products
}

func NewSimpleProductRepository(logger *log.Logger) *SimpleProductRepository {
	return &SimpleProductRepository{logger: logger, products: Products{}}
}

func (repo SimpleProductRepository) keys() []int {

	var keys []int

	for _, p := range repo.products {
		keys = append(keys, p.ID)
	}

	return keys
}

func (repo *SimpleProductRepository) GetAll() Products {
	repo.logger.Print("Récupération de tous les produits")
	return repo.products
}

func (repo *SimpleProductRepository) GetById(id int) Product {
	repo.logger.Printf("Récupération du produit ayant l'ID %d", id)

	var prod Product

	for _, p := range repo.products {
		if p.ID == id {
			prod = *p
			break
		}
	}
	return prod
}

func (repo *SimpleProductRepository) GetByName(name string) Product {
	repo.logger.Printf("Récupération du produit ayant le nom %s", name)

	var prod Product

	for _, p := range repo.products {
		if p.Name == name {
			prod = *p
			break
		}
	}
	return prod
}

func (repo *SimpleProductRepository) Create(product Product) Products {

	repo.logger.Printf("Création du produit : %#v ", product)

	var id int
	keys := repo.keys()

	if len(repo.products) == 0 {
		id = 1
	} else {
		sort.Ints(keys)
		lastKey := keys[len(keys)-1]
		id = lastKey + 1
	}
	product.ID = id
	product.CreationDate = time.Now().UTC().String()
	repo.products = append(repo.products, &product)

	return repo.products
}

func (repo *SimpleProductRepository) Update(product Product) Products {

	repo.logger.Printf("Mise à jour du produit : %#v", product)

	for _, p := range repo.products {
		if product.ID == (*p).ID {
			product.UpdateDate = time.Now().UTC().String()
			repo.products = append(repo.products, &product)
			break
		}
	}

	return repo.products
}

func (repo *SimpleProductRepository) Delete(id int) bool {

	repo.logger.Printf("Suppression du produit ayant l'ID %d", id)

	for i, p := range repo.products {
		if id == p.ID {

			p.RemovalDate = time.Now().UTC().String()
			repo.products = append(repo.products[:i], repo.products[(i+1):]...)
			return true
		}
	}

	return false
}
