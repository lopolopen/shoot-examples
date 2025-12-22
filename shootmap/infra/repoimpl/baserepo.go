package repoimpl

// import (
// 	"context"
// 	"shoot-examples/shootmap/domain/repo"

// 	"gorm.io/gorm"
// )

// type BaseRepo[TID any, TRoot any] struct {
// 	db *gorm.DB
// }

// func NewBaseRepo[TID any, TRoot any](db *gorm.DB) *BaseRepo[TID, TRoot] {
// 	r := &BaseRepo[TID, TRoot]{db: db}
// 	var _ repo.RepoBase[TID, TRoot] = r
// 	return r
// }

// func (r BaseRepo[TID, TRoot]) Get(ctx context.Context, id TID) (*TRoot, error) {

// }
