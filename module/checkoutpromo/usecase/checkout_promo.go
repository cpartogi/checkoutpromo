package usecase

import (
	"checkoutpromo/graph/model"
	"checkoutpromo/module/checkoutpromo/repository"
	"errors"
)

func ProductList() ([]*model.Product, error) {
	productList, err := repository.ProductList()

	if err != nil {
		return nil, err
	}
	return productList, nil
}

func CustomerList() ([]*model.Customer, error) {
	customerList, err := repository.CustomerList()

	if err != nil {
		return nil, err
	}
	return customerList, nil
}

func CartList(customerId string) ([]*model.Cart, error) {

	if customerId == "" {
		return nil, errors.New("customer_id must be filled")
	}

	cartList, err := repository.CartList(customerId)

	if err != nil {
		return nil, err
	}
	return cartList, nil
}

func OrderList(customerId string) ([]*model.Order, error) {

	if customerId == "" {
		return nil, errors.New("customer_id must be filled")
	}

	orderList, err := repository.OrderList(customerId)

	if err != nil {
		return nil, err
	}
	return orderList, nil
}

func OrderDetail(orderNum string) ([]*model.OrderDetail, error) {

	if orderNum == "" {
		return nil, errors.New("order_num must be filled")
	}

	orderDetail, err := repository.OrderDetail(orderNum)

	if err != nil {
		return nil, err
	}
	return orderDetail, nil
}
