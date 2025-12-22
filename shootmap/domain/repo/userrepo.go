package repo

import "shoot-examples/shootmap/domain/model"

type UserRepo interface {
	RepoBase[uint, model.User]
}
