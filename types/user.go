package types

import "time"

type RoleType string

// Role types
// 0 SuperAdmin; 1 Admin; 2 General

// User table model
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	RoleType  int8   `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
