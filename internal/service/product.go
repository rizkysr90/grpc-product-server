package service

import (
	"grpc-server-product/internal/store"
)

type ProductServicePort interface {
	Create(product *store.ProductData) error
	Update(pid uint64, updatedData *store.ProductData) error
	Delete(pid uint64) error
	GetAll() []store.ProductData
}
