package repository

import (
	"checkoutpromo/graph/model"
	"checkoutpromo/internal/db"

	uuid "github.com/nu7hatch/gouuid"
)

func AddCart(customerID string, productID string, qty int) (res *model.ResponseData, err error) {
	stmt, err := db.Db.Prepare(`INSERT INTO shopping_carts (cart_id, customer_id, product_id, qty, created_at) values (?,?,?,?, now()) `)

	if err != nil {
		return nil, err
	}

	u, _ := uuid.NewV4()
	Id := u.String()

	_, err = stmt.Exec(Id, customerID, productID, qty)

	if err != nil {
		return nil, err
	}

	hasil := &model.ResponseData{
		StatusCode: 200,
		Message:    "success add product to cart",
	}

	return hasil, err
}

func UpdateCart(customerID string, productID string, qty int) (res *model.ResponseData, err error) {
	stmt, err := db.Db.Prepare(`UPDATE shopping_carts set qty = ?, updated_at = now() WHERE customer_id= ? AND product_id = ? `)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(qty, customerID, productID)

	if err != nil {
		return nil, err
	}

	hasil := &model.ResponseData{
		StatusCode: 200,
		Message:    "success update product to cart",
	}

	return hasil, err
}

func ReduceStock(productID string, qty int) (err error) {
	stmt, err := db.Db.Prepare(`UPDATE products set product_qty=product_qty-? WHERE product_id = ?`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(qty, productID)

	if err != nil {
		return err
	}

	return err
}

func AddStock(productID string, qty int) (err error) {
	stmt, err := db.Db.Prepare(`UPDATE products set product_qty=product_qty+? WHERE product_id = ?`)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(qty, productID)

	if err != nil {
		return err
	}

	return err
}

func DeleteCart(customerID string, productID string) (res *model.ResponseData, err error) {
	stmt, err := db.Db.Prepare(`DELETE shopping_carts WHERE customer_id = ? AND product_id = ?`)

	if err != nil {
		return nil, err
	}

	u, _ := uuid.NewV4()
	Id := u.String()

	_, err = stmt.Exec(Id, customerID, productID)

	if err != nil {
		return nil, err
	}

	hasil := &model.ResponseData{
		StatusCode: 200,
		Message:    "success remove product from cart",
	}

	return hasil, err
}
