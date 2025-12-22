package po

import (
	"shoot-examples/shootmap/domain/enums"
	"shoot-examples/shootmap/domain/model"
	"shoot-examples/shootmap/infra/mapper"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Order struct {
	mapper.SQLMapper
	*gorm.Model
	ID        string            `gorm:"size:32;primarykey"`
	Amount    decimal.Decimal   `gorm:"type:decimal(20,2)"`
	Status    enums.OrderStatus `gorm:"type:enum('Pending', 'Completed', 'Canceled')"`
	City      string            `gorm:"size:64"`
	Street    string            `gorm:"size:64"`
	Room      string            `gorm:"size:64"`
	OrderTime time.Time
}

func (po *Order) writeDomain(dm *model.Order) {
	dm.Address = model.OrderAddress{
		City:   po.City,
		Street: po.Street,
		Room:   po.Room,
	}
}

func (po *Order) readDomain(dm model.Order) {
	po.Model = nil
	po.City = dm.Address.City
	po.Street = dm.Address.Street
	po.Room = dm.Address.Room
}
