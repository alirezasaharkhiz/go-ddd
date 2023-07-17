package customer

import (
	"ddd/aggregator"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("customer not found")
	ErrFailedToAddCustomer    = errors.New("failed to create customer")
	ErrFailedToUpdateCustomer = errors.New("failed to update customer")
)

type Repository interface {
	Get(uuid uuid.UUID) (aggregator.Customer, error)
	Add(aggregator.Customer) error
	Update(aggregator.Customer) error
}
