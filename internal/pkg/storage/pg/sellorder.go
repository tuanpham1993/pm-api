package pg

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SellOrder struct {
	ID         string `gorm:"type:uuid;primary_key;" json:"id"`
	CustomerID *string `json:"customerId"`
	Customer   *Customer
	Items      []SellOrderItem `gorm: "foreignKey:SellOrderID"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
}

func (sellOrder *SellOrder) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}
