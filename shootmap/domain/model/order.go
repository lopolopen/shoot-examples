package model

import (
	"shoot-examples/shootmap/domain/enums"
	"time"

	"github.com/shopspring/decimal"
)

type Order struct {
	ID        string
	Amount    decimal.Decimal
	Status    enums.OrderStatus
	OrderTime time.Time
	Address   OrderAddress
}
