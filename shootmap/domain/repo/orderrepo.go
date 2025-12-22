package repo

import "shoot-examples/shootmap/domain/model"

type OrderRepo interface {
	RepoBase[string, model.Order]
}
