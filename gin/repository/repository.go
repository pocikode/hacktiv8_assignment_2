package repository

type OrderRepository interface {
	CreateOrder(order *Order) error
	GetOrders() (*[]Order, error)
	GetOrderByID(id int) (*Order, error)
	UpdateOrder(order *Order) error
	DeleteOrder(order *Order) error
}
