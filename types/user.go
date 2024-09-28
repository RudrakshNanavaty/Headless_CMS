package types

import "time"

// User table model
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	RoleType  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
