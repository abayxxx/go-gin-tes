package service

import "go-gin/app/domain/dto"

type ProductService interface {
	GetAll() (response []dto.ProductResponse, err error)
	GetById(id uint64) (response dto.ProductResponse, err error)
	StoreProduct(request dto.ProductRequest) (response bool, err error)
	UpdateProduct(id uint64, request dto.UpdateProductRequest) (response bool, err error)
	DeleteProduct(id uint64) (response bool, err error)
}
