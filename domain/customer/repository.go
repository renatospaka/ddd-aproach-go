package customer

import (
	"errors"

	"github.com/google/uuid"
	"github.com/renatospaka/tavern/aggregate"
)

var (
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer = errors.New("failed to update the customer in the repository")
)

// CustomerRepository is a interface that defines the rules around what a customer repository
// Has to be able to perform
type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}