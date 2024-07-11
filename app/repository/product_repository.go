package repository

import (
	"go-gin/app/domain/dto"
	"go-gin/app/domain/model"
)

type ProductRepository interface {
	GetAll() (response []model.Product, err error)
	GetById(id uint64) (response model.Product, err error)
	StoreProduct(product dto.ProductRequest) (response bool, err error)
	UpdateProduct(productId uint64, product dto.UpdateProductRequest) (response bool, err error)
	DeleteProduct(id uint64) (response bool, err error)
}
