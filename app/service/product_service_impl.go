package service

import (
	"go-gin/app/domain/dto"
	"go-gin/app/repository"
	"go-gin/app/service/jwt"
)

type ProductServiceImpl struct {
	repository repository.ProductRepository
	JwtService jwt.JwtService
}

// NewProductServiceImpl is a constructor function to create authServiceImpl
func NewProductServiceImpl(repository repository.ProductRepository, jwtService jwt.JwtService) *ProductServiceImpl {
	return &ProductServiceImpl{repository, jwtService}
}

func (s *ProductServiceImpl) GetAll() (response []dto.ProductResponse, err error) {
	resp, err := s.repository.GetAll()
	if err != nil {
		return response, err
	}

	//convert to dto
	response = make([]dto.ProductResponse, len(resp))
	for i, product := range resp {
		response[i] = dto.ProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Photo:       product.Photo,
			Stock:       product.Stock,
		}
	}

	return response, nil
}

func (s *ProductServiceImpl) GetById(id uint64) (response dto.ProductResponse, err error) {
	resp, err := s.repository.GetById(id)
	if err != nil {
		return response, err
	}

	//convert to dto
	response = dto.ProductResponse{
		ID:          resp.ID,
		Name:        resp.Name,
		Price:       resp.Price,
		Description: resp.Description,
		Photo:       resp.Photo,
		Stock:       resp.Stock,
	}

	return response, nil
}

func (s *ProductServiceImpl) StoreProduct(request dto.ProductRequest) (response bool, err error) {
	if err = validate.Struct(request); err != nil {
		return response, err
	}

	resp, err := s.repository.StoreProduct(request)
	if err != nil {
		return false, err
	}

	return resp, nil
}

func (s *ProductServiceImpl) UpdateProduct(id uint64, request dto.UpdateProductRequest) (response bool, err error) {
	if err = validate.Struct(request); err != nil {
		return response, err
	}

	resp, err := s.repository.UpdateProduct(id, request)
	if err != nil {
		return response, err
	}

	return resp, nil
}

func (s *ProductServiceImpl) DeleteProduct(id uint64) (response bool, err error) {
	resp, err := s.repository.DeleteProduct(id)
	if err != nil {
		return response, err
	}

	return resp, nil
}
