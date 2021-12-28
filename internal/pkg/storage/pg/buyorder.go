package pg

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BuyOrder struct {
	ID         string `gorm:"type:uuid;primary_key;" json:"id"`
	SupplierID string `json:"supplierId"`
	Supplier   Supplier
	Items      []BuyOrderItem `gorm: "foreignKey:BuyOrderID"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
}

func (buyOrder *BuyOrder) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}
