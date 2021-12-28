package pg

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type SellOrderItem struct {
	ID          string `gorm:"type:uuid;primary_key;" json:"id"`
	SellOrderID string `json:"sellOrderId"`
	Product     Product
	ProductID   string    `json:"productId"`
	Price       int64     `json:"price"`
	Quantity    int64     `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (sellOrderItem *SellOrderItem) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}
