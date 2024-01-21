package product

import (
	"context"
	productService "grpc-server-product/internal/service/product"
	"grpc-server-product/internal/store"
	"io"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/rizkysr90/my-protobuf/gen/go/personal/productservice/product"
)

type ProductGRPC struct {
	*product.UnimplementedProductServiceServer
	service *productService.ProductService
}

func New(service *productService.ProductService) *ProductGRPC {
	return &ProductGRPC{
		service: service,
	}
}

func (p *ProductGRPC) Create(ctx context.Context, req *product.Product) (*product.Response, error) {
	creationProduct := store.ProductData{
		Pid:   req.GetPid(),
		Name:  req.GetName(),
		Stock: uint64(req.GetStock()),
		Price: req.GetPrice(),
	}
	if err := p.service.Create(&creationProduct); err != nil {
		return &product.Response{Status: uint32(0)}, nil
	}
	return &product.Response{Status: uint32(1)}, nil
}
func (p *ProductGRPC) Update(ctx context.Context, req *product.UpdateProduct) (*product.Response, error) {
	creationProduct := store.ProductData{
		Pid:   req.UpdatedDataProduct.GetPid(),
		Name:  req.UpdatedDataProduct.GetName(),
		Stock: uint64(req.UpdatedDataProduct.GetStock()),
		Price: req.UpdatedDataProduct.GetPrice(),
	}
	if err := p.service.Update(uint64(req.GetPid()), &creationProduct); err != nil {
		return &product.Response{Status: uint32(0)}, nil
	}
	return &product.Response{Status: uint32(1)}, nil
}
func (p *ProductGRPC) Delete(ctx context.Context, req *product.DeleteProduct) (*product.Response, error) {
	if err := p.service.Delete(uint64(req.GetPid())); err != nil {
		return &product.Response{Status: uint32(0)}, nil
	}
	return &product.Response{Status: uint32(1)}, nil
}
func (p *ProductGRPC) GetList(context.Context, *product.ListProduct) (*product.Response, error) {
	return nil, nil
}
func (p *ProductGRPC) CreateProducts(stream product.ProductService_CreateProductsServer) error {
	for {
		productInput, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		data := store.ProductData{
			Pid:   productInput.Pid,
			Name:  productInput.Name,
			Stock: uint64(productInput.Stock),
			Price: productInput.Price,
		}
		p.service.Create(&data)
	}
	return stream.SendAndClose(&product.Response{Status: uint32(1)})
}
func (p *ProductGRPC) GetListStream(_ *empty.Empty, stream product.ProductService_GetListStreamServer) error {
	getProducts := p.service.GetAll()
	for idx, item := range getProducts {
		toProduct := &product.Product{
			Pid:   item.Pid,
			Name:  item.Name,
			Stock: int64(item.Stock),
			Price: item.Price,
		}
		if err := stream.Send(toProduct); err != nil {
			return err
		}
		log.Println("Sent value ", idx+1)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}
