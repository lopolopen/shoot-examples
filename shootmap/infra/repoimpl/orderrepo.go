package repoimpl

import (
	"context"
	"shoot-examples/shootmap/domain/model"
	"shoot-examples/shootmap/domain/qpo"
	"shoot-examples/shootmap/domain/repo"
	"shoot-examples/shootmap/infra/po"

	"gorm.io/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

// Get implements [repo.OrderRepo].
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

// Query implements [repo.OrderRepo].
func (r *OrderRepo) Query(ctx context.Context, page qpo.Pagination) ([]*model.Order, int64, error) {
	var orderPOs []*po.Order
	var total int64

	if err := r.db.WithContext(ctx).
		Model(&po.Order{}).
		Count(&total).Error; err != nil {
		return nil, 0, err
	}
	err := r.db.WithContext(ctx).
		Model(&po.Order{}).
		Offset(page.Offset()).
		Limit(page.PerPage).
		Scan(&orderPOs).Error
	if err != nil {
		return nil, 0, err
	}

	orders := make([]*model.Order, len(orderPOs))
	for i := range orderPOs {
		orders[i] = orderPOs[i].ToDomain()
	}
	return orders, total, nil
}

func NewOrderRepo(db *gorm.DB) *OrderRepo {
	r := &OrderRepo{db: db}
	var _ repo.OrderRepo = r
	return r
}
