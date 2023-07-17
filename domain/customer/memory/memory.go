package memory

import (
	"ddd/aggregator"
	"ddd/domain/customer"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregator.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregator.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregator.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}

	return aggregator.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregator.Customer) error {
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

func (mr *MemoryRepository) Update(c aggregator.Customer) error {
	if _, ok := mr.customers[c.GetId()]; !ok {
		return fmt.Errorf("customer not found: %w", customer.ErrFailedToUpdateCustomer)
	}

	mr.Lock()
	mr.customers[c.GetId()] = c

	return nil
}
