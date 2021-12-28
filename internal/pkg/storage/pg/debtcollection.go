package pg

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type DebtCollection struct {
	ID         string `gorm:"type:uuid;primary_key;" json:"id"`
	CustomerID string `json:"customerId"`
	Customer   Customer
	Amount     int64     `json:"amount"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (payment *DebtCollection) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}
