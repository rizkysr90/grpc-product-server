package product

import (
	"errors"
	store "grpc-server-product/internal/store"
	"log"
	"slices"
)

type ProductService struct {
	productData []store.ProductData
}

func New(masterData []store.ProductData) *ProductService {
	return &ProductService{
		productData: masterData,
	}
}
func (p *ProductService) Create(product *store.ProductData) error {
	p.productData = append(p.productData, *product)
	log.Println(p.productData)
	return nil
}

// func (p *ProductService) GetList()
func (p *ProductService) Update(pid uint64, updatedData *store.ProductData) error {
	getIdxData := -1
	for idx, product := range p.productData {
		if product.Pid == pid {
			getIdxData = idx
			break
		}
	}
	if getIdxData == -1 {
		return errors.New("data not found")
	}
	p.productData[getIdxData].Name = updatedData.Name
	p.productData[getIdxData].Price = updatedData.Price
	p.productData[getIdxData].Stock = updatedData.Stock
	log.Println(p.productData)

	return nil
}
func (p *ProductService) Delete(pid uint64) error {
	getIdxData := -1
	for idx, product := range p.productData {
		if product.Pid == pid {
			getIdxData = idx
			break
		}
	}
	if getIdxData == -1 {
		return errors.New("data not found")
	}
	p.productData = slices.Delete(p.productData, getIdxData, getIdxData+1)
	log.Println(p.productData)
	return nil
}
func (p *ProductService) GetAll() []store.ProductData {
	return p.productData
}
