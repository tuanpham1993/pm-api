package pg

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	ID        string    `gorm:"type:uuid;primary_key;" json:"id"`
	Name      string    `json:"name"`
	Debt      int64     `json:"debt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (supplier *Supplier) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}
