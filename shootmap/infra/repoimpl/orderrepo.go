package repoimpl

import (
	"context"
	"shoot-examples/shootmap/domain/model"
	"shoot-examples/shootmap/domain/repo"
	"shoot-examples/shootmap/infra/po"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	r := &OrderRepo{db: db}
	var _ repo.OrderRepo = r
	return r
}

func (r *OrderRepo) Get(ctx context.Context, id string) (*model.Order, error) {
	var orderPO po.Order
	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&orderPO).Error
	if err != nil {
		return nil, err
	}

	order := orderPO.ToDomain()
	return order, nil
}
