package service

import (
	"fmt"
	"go-gin/app/domain/dto"
	helper "go-gin/app/helper"
	"go-gin/app/repository"
	"go-gin/app/service/jwt"
)

type OrderServiceImpl struct {
	repository repository.OrderRepository
	authRepo   repository.AuthRepository
	JwtService jwt.JwtService
}

// NewOrderServiceImpl is a constructor function to create NewOrderServiceImpl
func NewOrderServiceImpl(repository repository.OrderRepository, authRepo repository.AuthRepository, jwtService jwt.JwtService) *OrderServiceImpl {
	return &OrderServiceImpl{repository, authRepo, jwtService}
}

func (s *OrderServiceImpl) GetAllOrderUser(userId uint64) (response []dto.OrderHistoryResponse, err error) {
	resp, err := s.repository.GetAllOrderUser(userId)
	if err != nil {
		return response, err
	}

	//convert to dto
	response = make([]dto.OrderHistoryResponse, len(resp))
	for i, order := range resp {
		total := order.Price * float64(order.Quantity)
		var status string
		if order.Status == 1 {
			status = "Pending"
		} else if order.Status == 2 {
			status = "Success"
		} else {
			status = "Failed"
		}
		response[i] = dto.OrderHistoryResponse{
			ID:          uint64(order.ID),
			ProductID:   uint64(order.ProductID),
			ProductName: order.Product.Name,
			Price:       order.Price,
			Quantity:    order.Quantity,
			Total:       total,
			Status:      status,
		}
	}

	return response, nil
}

func (s *OrderServiceImpl) GetDetailOrderUser(userId, orderId uint64) (response dto.OrderHistoryResponse, err error) {
	resp, err := s.repository.GetDetailOrderUser(userId, orderId)
	if err != nil {
		return response, err
	}

	//convert to dto
	total := resp.Price * float64(resp.Quantity)
	var status string
	if resp.Status == 1 {
		status = "Pending"
	} else if resp.Status == 2 {
		status = "Success"
	} else {
		status = "Failed"
	}

	response = dto.OrderHistoryResponse{
		ID:          uint64(resp.ID),
		ProductID:   uint64(resp.ProductID),
		ProductName: resp.Product.Name,
		Price:       resp.Price,
		Quantity:    resp.Quantity,
		Total:       total,
		Status:      status,
	}

	return response, nil
}

func (s *OrderServiceImpl) StoreShoppingCart(userId uint64, request dto.ShoppingCartRequest) (response bool, err error) {
	if err = validate.Struct(request); err != nil {
		return response, err
	}

	response, err = s.repository.StoreShoppingCart(userId, request)
	if err != nil {
		return response, err
	}

	return true, nil
}

func (s *OrderServiceImpl) GetShoppingCartList(userId uint64) (response []dto.ShoppingCartResponse, err error) {
	resp, err := s.repository.GetShoppingCartList(userId)
	if err != nil {
		return response, err
	}

	//convert to dto
	response = make([]dto.ShoppingCartResponse, len(resp))
	for i, cart := range resp {
		total := cart.Product.Price * float64(cart.Quantity)
		response[i] = dto.ShoppingCartResponse{
			ID:                 uint64(cart.ID),
			ProductID:          uint64(cart.ProductID),
			ProductName:        cart.Product.Name,
			ProductPhoto:       cart.Product.Photo,
			ProductDescription: cart.Product.Description,
			Price:              cart.Product.Price,
			Quantity:           cart.Quantity,
			Total:              total,
		}
	}

	return response, nil
}

func (s *OrderServiceImpl) StoreOrder(userId uint64, request dto.OrderRequest) (response bool, err error) {
	if err = validate.Struct(request); err != nil {
		return response, err
	}

	response, err = s.repository.StoreOrder(userId, request)
	if err != nil {
		return response, err
	}

	//get user by id
	user, err := s.authRepo.GetUserById(uint(userId))
	if err != nil {
		return response, err
	}

	//send email
	fmt.Println("Sending email to:", user.Email)
	err = helper.SendEmail(user.Email, "Order Confirmation")
	if err != nil {
		fmt.Println("Error sending email:", err)
	}

	return true, nil
}
