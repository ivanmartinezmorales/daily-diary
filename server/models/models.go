package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GenerateISOString() string {
	return time.Now().UTC().Format("2020-01-02T11:11:05999Z07:00")
}

type Base struct {
	ID        uint      `gorm:"primaryKey"`
	UUID      uuid.UUID `json:"_id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.UUID = uuid.New()
	t := GenerateISOString()

	b.CreatedAt, b.UpdatedAt = t, t

	return nil
}

func (b *Base) AfterUpdate(tx *gorm.DB) error {
	b.UpdatedAt = GenerateISOString()
	return nil
}
