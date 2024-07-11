package repository

import (
	"errors"
	"fmt"
	"go-gin/app/domain/dto"
	"go-gin/app/domain/model"
	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

// NewOrderRepositoryImpl is a constructor function to create NewOrderRepositoryImpl
func NewOrderRepositoryImpl(db *gorm.DB) *OrderRepositoryImpl {
	return &OrderRepositoryImpl{db}
}

func (m *OrderRepositoryImpl) GetAllOrderUser(userId uint64) (response []model.UserOrderJournal, err error) {
	var orders []model.UserOrderJournal

	fmt.Println("userId", userId)
	err = m.db.Table("user_order_journals").Where("user_id = ?", userId).
		Preload("Product").
		Find(&orders).Error
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func (m *OrderRepositoryImpl) GetDetailOrderUser(userId uint64, orderId uint64) (response model.UserOrderJournal, err error) {
	var order model.UserOrderJournal
	err = m.db.Table("user_order_journals").Where("user_id = ? AND id = ?", userId, orderId).
		Preload("Product").
		First(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}

func (m *OrderRepositoryImpl) StoreShoppingCart(userId uint64, request dto.ShoppingCartRequest) (response bool, err error) {

	// check product exist
	var product model.Product
	err = m.db.Table("products").Where("id = ?", request.ProductID).First(&product).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("Product with id %d not found", request.ProductID))
	}

	//check stock
	if product.Stock < request.Quantity {
		return false, errors.New(fmt.Sprintf("Stock is not enough"))
	}

	fmt.Println("userId", userId)

	cart := model.ShoppingCart{
		UserID:    uint(userId),
		ProductID: uint(request.ProductID),
		Quantity:  request.Quantity,
	}

	err = m.db.Table("shopping_carts").Create(&cart).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

func (m *OrderRepositoryImpl) GetShoppingCartList(userId uint64) (response []model.ShoppingCart, err error) {
	var carts []model.ShoppingCart
	err = m.db.Table("shopping_carts").Where("user_id = ?", userId).
		Preload("Product").
		Find(&carts).Error
	if err != nil {
		return carts, err
	}
	return carts, nil
}

func (m *OrderRepositoryImpl) StoreOrder(userId uint64, request dto.OrderRequest) (response bool, err error) {

	// get product price
	var product model.Product
	err = m.db.Table("products").Where("id = ?", request.ProductID).First(&product).Error
	if err != nil {
		return false, err
	}

	//get balance user
	var balance model.UserWallet
	err = m.db.Table("user_wallets").Where("user_id = ?", userId).First(&balance).Error
	if err != nil {
		return false, errors.New(fmt.Sprintf("Wallet User with id %d not found", userId))
	}

	// check balance
	if balance.Balance < product.Price {
		return false, errors.New(fmt.Sprintf("Your balance is not enough, please top up your balance"))
	}

	//check stock
	if product.Stock < request.Quantity {
		return false, errors.New(fmt.Sprintf("Stock is not enough"))
	}

	//decrease balance
	balance.Balance = balance.Balance - product.Price
	err = m.db.Table("user_wallets").Where("user_id = ?", userId).Updates(&balance).Error
	if err != nil {
		return false, err
	}

	//store order
	price := product.Price * float64(request.Quantity)
	order := model.UserOrderJournal{
		UserID:          uint(userId),
		ProductID:       uint(request.ProductID),
		Quantity:        request.Quantity,
		Price:           price,
		Status:          2,
		LastBalanceUser: balance.Balance,
	}

	err = m.db.Table("user_order_journals").Create(&order).Error
	if err != nil {
		return false, err
	}

	//decrease stock
	product.Stock = product.Stock - request.Quantity
	err = m.db.Table("products").Select("stocks").Where("id = ?", request.ProductID).Updates(&product).Error
	if err != nil {
		return false, err
	}

	//delete shopping cart with related product
	err = m.db.Table("shopping_carts").Where("user_id = ? AND product_id = ?", userId, request.ProductID).Delete(&model.ShoppingCart{}).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
