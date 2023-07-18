package product

import (
	"ddd/aggregator"
	"github.com/google/uuid"
)

type Repository interface {
	GetAll() ([]aggregator.Product, error)
	GetByID(id uuid.UUID) (aggregator.Product, error)
	Add(product aggregator.Product) error
	Update(product aggregator.Product) error
	Delete(id uuid.UUID) error
}
