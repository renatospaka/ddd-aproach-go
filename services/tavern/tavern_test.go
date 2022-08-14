package tavern

import (
	"testing"

	"github.com/google/uuid"
	"github.com/renatospaka/tavern/domain/product"
	"github.com/renatospaka/tavern/services/order"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.99)
	if err != nil {
		t.Error(err)
	}

	peenuts, err := product.NewProduct("Peenuts", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}

	wine, err := product.NewProduct("Wine", "Healthy Snacks", 0.99)
	if err != nil {
		t.Error(err)
	}

	products := []product.Product{
		beer, peenuts, wine,
	}
	return products
}

func Test_Tavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	orderServ, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(orderServ))
	if err != nil {
		t.Error(err)
	}

	uid, err := orderServ.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}
}

func Test_MongoTavern(t *testing.T) {
	// Create OrderService
	products := init_products(t)

	orderServ, err := order.NewOrderService(
		order.WithMongoCustomerRepository("mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavern(WithOrderService(orderServ))
	if err != nil {
		t.Error(err)
	}

	uid, err := orderServ.AddCustomer("Percy")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		t.Error(err)
	}
}
