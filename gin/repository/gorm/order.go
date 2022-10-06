package gorm

import (
	"assignment2/gin/repository"
	"gorm.io/gorm"
)

type OrderGorm struct {
	db *gorm.DB
}

func (o *OrderGorm) CreateOrder(order *repository.Order) error {
	return o.db.Create(order).Error
}

func (o *OrderGorm) GetOrders() (*[]repository.Order, error) {
	var orders []repository.Order

	err := o.db.Preload("Items").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return &orders, nil
}

func (o *OrderGorm) GetOrderByID(id int) (*repository.Order, error) {
	var order repository.Order

	err := o.db.First(&order, id).Error
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (o *OrderGorm) UpdateOrder(order *repository.Order) error {
	return o.db.Save(&order).Error
}

func (o *OrderGorm) DeleteOrder(order *repository.Order) error {
	return o.db.Delete(&order).Error
}

func NewOrderRepo(db *gorm.DB) repository.OrderRepository {
	return &OrderGorm{
		db: db,
	}
}
