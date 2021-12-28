package pg

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID         string    `gorm:"type:uuid;primary_key;" json:"id"`
	Name       string    `json:"name"`
	Unit       string    `json:"unit"`
	Quantity   int64     `json:"quantity"`
	InputPrice int64     `json:"inputPrice"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func (product *Product) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}
