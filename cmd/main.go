package main

import (
	"github.com/google/uuid"
	"github.com/renatospaka/tavern/domain/product"
	"github.com/renatospaka/tavern/services/order"
	servicetavern "github.com/renatospaka/tavern/services/tavern"
)

func main() {

	products := productInventory()
	// Create Order Service to use in tavern
	os, err := order.NewOrderService(
		order.WithMongoCustomerRepository("mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}
	
	// Create tavern service
	tavern, err := servicetavern.NewTavern(
		servicetavern.WithOrderService(os))
	if err != nil {
		panic(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}
	
	order := []uuid.UUID{
		products[0].GetID(),
	}
	// Execute Order
	err = tavern.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		panic(err)
	}
	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		panic(err)
	}
	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}