package repo

import (
	"context"
	"shoot-examples/shootmap/domain/qpo"
)

type RepoBase[TID any, TEntity any] interface {
	Get(ctx context.Context, id TID) (*TEntity, error)

	Query(ctx context.Context, page qpo.Pagination) ([]*TEntity, int64, error)
}
