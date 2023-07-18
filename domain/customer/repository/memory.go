package repository

import (
	"ddd/aggregator"
	"ddd/domain/customer"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type CustomerMemoryRepository struct {
	customers map[uuid.UUID]aggregator.Customer
	sync.Mutex
}

func NewCustomerMemoryRepository() *CustomerMemoryRepository {
	return &CustomerMemoryRepository{
		customers: make(map[uuid.UUID]aggregator.Customer),
	}
}

func (mr *CustomerMemoryRepository) Get(id uuid.UUID) (aggregator.Customer, error) {
	if cs, ok := mr.customers[id]; ok {
		return cs, nil
	}

	return aggregator.Customer{}, customer.ErrCustomerNotFound
}

func (mr *CustomerMemoryRepository) Add(c aggregator.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregator.Customer)
		mr.Unlock()
	}

	if _, ok := mr.customers[c.GetId()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrFailedToAddCustomer)
	}

	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()

	return nil
}

func (mr *CustomerMemoryRepository) Update(c aggregator.Customer) error {
	if _, ok := mr.customers[c.GetId()]; !ok {
		return fmt.Errorf("customer not found: %w", customer.ErrFailedToUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetId()] = c

	return nil
}
