// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Cart struct {
	CartID      string  `json:"cart_id"`
	CustomerID  string  `json:"customer_id"`
	ProductID   string  `json:"product_id"`
	Sku         string  `json:"sku"`
	ProductName string  `json:"product_name"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
	CreatedAt   string  `json:"created_at"`
}

type Customer struct {
	CustomerID    string `json:"customer_id"`
	CustomerName  string `json:"customer_name"`
	CustomerEmail string `json:"customer_email"`
}

type Order struct {
	OrderNum   string `json:"order_num"`
	CustomerID string `json:"customer_id"`
}

type OrderDetail struct {
	OrderID     string  `json:"order_id"`
	OrderNum    string  `json:"order_num"`
	CustomerID  string  `json:"customer_id"`
	ProductID   string  `json:"product_id"`
	Sku         string  `json:"sku"`
	ProductName string  `json:"product_name"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int     `json:"quantity"`
	TotalPrice  float64 `json:"total_price"`
}

type Product struct {
	ProductID    string  `json:"product_id"`
	Sku          string  `json:"sku"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	ProductQty   int     `json:"product_qty"`
}

type ResponseData struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
