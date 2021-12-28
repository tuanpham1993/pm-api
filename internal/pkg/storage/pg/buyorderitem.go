package pg

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BuyOrderItem struct {
	ID         string `gorm:"type:uuid;primary_key;" json:"id"`
	BuyOrderID string `json:"buyOrderId"`
	Product    Product
	ProductID  string    `json:"productId"`
	Price      int64     `json:"price"`
	Quantity   int64     `json:"quantity"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (buyOrderItem *BuyOrderItem) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}
