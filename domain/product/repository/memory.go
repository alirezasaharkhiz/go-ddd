package repository

import (
	"ddd/aggregator"
	"github.com/google/uuid"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregator.Product
	sync.Mutex
}

func NewMemoryProductRepository() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregator.Product),
	}
}

//Add(product aggregator.Product) error
//Update(product aggregator.Product) error
//Delete(id uuid.UUID) error

func (mpr *MemoryProductRepository) GetAll() ([]aggregator.Product, error) {
	var products []aggregator.Product
	for _, p := range mpr.products {
		products = append(products, p)
	}

	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregator.Product, error) {
	return mpr.products[id], nil
}

func (mpr *MemoryProductRepository) Add(product aggregator.Product) error {
	mpr.products[uuid.New()] = product
	return nil
}

func (mpr *MemoryProductRepository) Update(product aggregator.Product) error {
	mpr.products[product.GetID()] = product
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	delete(mpr.products, id)
	return nil
}
