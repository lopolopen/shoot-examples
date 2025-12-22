package repoimpl

import (
	"shoot-examples/shootmap/infra/po"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	dests := []any{&po.Order{}, &po.User{}, &po.UserAddress{}}
	for _, d := range dests {
		err := db.AutoMigrate(d)
		if err != nil {
			return err
		}
	}
	return nil
}
