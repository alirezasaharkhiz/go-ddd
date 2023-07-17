package services

import (
	"ddd/domain/customer"
	"ddd/domain/customer/repository"
)

type OrderConfig func(os *OrderService) error

type OrderService struct {
	customer.Repository
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
		os.Repository = cr
		return nil
	}
}

func WithCustomerMemoryRepository() OrderConfig {
	cmr := repository.NewCustomerMemoryRepository()
	return WithCustomerRepository(cmr)
}
