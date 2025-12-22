package enums

import (
	"database/sql/driver"
	"errors"

	"github.com/lopolopen/shoot"
)

//go:generate go tool shoot enum -json -text -file=$GOFILE

type OrderStatus int32

const (
	OrderStatusPending OrderStatus = iota
	OrderStatusCompleted
	OrderStatusCanceled
)

func (s OrderStatus) Value() (driver.Value, error) {
	return s.String(), nil
}

func (s *OrderStatus) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New("bad status type")
	}
	e, err := shoot.ParseEnum[OrderStatus](string(data))
	if err != nil {
		return err
	}
	*s = e
	return nil
}
