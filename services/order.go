package services

import (
	"ddd/domain/customer"
	"ddd/domain/customer/repository"
	"fmt"
	"github.com/google/uuid"
)

type OrderConfig func(os *OrderService) error

type OrderService struct {
	customerRepo customer.Repository
}

func NewOrderService(cfgs ...OrderConfig) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// WithCustomerRepository returns a function
func WithCustomerRepository(cr customer.Repository) OrderConfig {
	return func(os *OrderService) error {
		os.customerRepo = cr
		return nil
	}
}

func WithCustomerMemoryRepository() OrderConfig {
	cmr := repository.NewCustomerMemoryRepository()
	return WithCustomerRepository(cmr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productIDs []uuid.UUID) error {
	c, err := o.customerRepo.Get(customerID)
	if err != nil {
		return err
	}

	fmt.Println(c)

	return nil
}
