package test

import (
	"fmt"
	"github.com/hsedjame/products-api/client/client"
	"github.com/hsedjame/products-api/client/client/products"
	"testing"
)

func TestClient(t *testing.T) {
	config := client.DefaultTransportConfig().WithHost("localhost:8080")
	cli := client.NewHTTPClientWithConfig(nil, config)
	prods, err := cli.Products.ListOfProducts(products.NewListOfProductsParams(), nil)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(prods.Payload)
}
