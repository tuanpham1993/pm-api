package pg

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID        string    `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string    `json:"name"`
	Debt      int64     `json:"debt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (customer *Customer) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}
