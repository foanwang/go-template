package model

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primaryKey,autoIncrement"`
	Name      string
	Email     *string `gorm:"uniqueIndex"`
	Age       uint8
	Birthday  *time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
}

func (b *User) TableName() string {
	return "user"
}
