package repository

import (
	"go-gin/app/domain/dto"
	"go-gin/app/domain/model"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	db *gorm.DB
}

// NewProductRepositoryImpl is a constructor function to create NewProductRepositoryImpl
func NewProductRepositoryImpl(db *gorm.DB) *ProductRepositoryImpl {
	return &ProductRepositoryImpl{db}
}

func (m *ProductRepositoryImpl) GetAll() (response []model.Product, err error) {
	var products []model.Product
	err = m.db.Table("products").Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (m *ProductRepositoryImpl) GetById(id uint64) (response model.Product, err error) {
	product := model.Product{}
	err = m.db.Table("products").Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (m *ProductRepositoryImpl) StoreProduct(product dto.ProductRequest) (response bool, err error) {
	err = m.db.Model(&model.Product{}).Table("products").Create(&product).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *ProductRepositoryImpl) UpdateProduct(productId uint64, product dto.UpdateProductRequest) (response bool, err error) {
	err = m.db.Table("products").Where("id = ?", productId).Updates(&product).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (m *ProductRepositoryImpl) DeleteProduct(id uint64) (response bool, err error) {
	err = m.db.Table("products").Where("id = ?", id).Delete(&model.Product{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
