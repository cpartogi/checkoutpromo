package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"checkoutpromo/graph/generated"
	"checkoutpromo/graph/model"
	"checkoutpromo/module/checkoutpromo/usecase"
	"context"
	"fmt"
)

func (r *mutationResolver) AddCart(ctx context.Context, input model.AddCart) (*model.ResponseData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteCart(ctx context.Context, input model.DeleteCart) (*model.ResponseData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Checkout(ctx context.Context, input model.Checkout) (*model.ResponseData, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ProductList(ctx context.Context) ([]*model.Product, error) {
	productList, err := usecase.ProductList()

	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (r *queryResolver) CustomerList(ctx context.Context) ([]*model.Customer, error) {
	customerList, err := usecase.CustomerList()

	if err != nil {
		return nil, err
	}

	return customerList, nil
}

func (r *queryResolver) ShoppingCart(ctx context.Context, customerID string) ([]*model.Cart, error) {
	cartList, err := usecase.CartList(customerID)

	if err != nil {
		return nil, err
	}

	return cartList, nil
}

func (r *queryResolver) OrderByCustomer(ctx context.Context, customerID string) ([]*model.Order, error) {
	orderList, err := usecase.OrderList(customerID)

	if err != nil {
		return nil, err
	}

	return orderList, nil
}

func (r *queryResolver) OrderDetail(ctx context.Context, orderNum string) ([]*model.OrderDetail, error) {

	orderDetail, err := usecase.OrderDetail(orderNum)

	if err != nil {
		return nil, err
	}

	return orderDetail, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
