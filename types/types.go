package types

import (
	"time"
)

// Type table model
type Type struct {
	ID   uint   `json:"type_id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null"`
}

// Content table model
type Content struct {
	ID         uint `json:"content_id" gorm:"primaryKey"`
	TypeID     uint `json:"type_id" gorm:"not null"`
	Type       Type `gorm:"foreignKey:TypeID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Data       Data        `json:"data" gorm:"foreignKey:ContentID"`
	Attributes []Attribute `json:"attributes" gorm:"foreignKey:ContentID"`
	Children   []Content   `json:"children" gorm:"many2many:children;"`
}

// Data table model
type Data struct {
	ID        uint     `json:"data_id" gorm:"primaryKey"`
	ContentID uint     `json:"content_id" gorm:"not null"`
	Text      []string `json:"text" gorm:"type:text[]"`
	ImageUrls []string `json:"image" gorm:"type:text[]"`
	PdfUrls   []string `json:"pdf" gorm:"type:text[]"`
	Code      []string `json:"code" gorm:"type:text[]"`
	CreatedAt time.Time
}

// Attribute table model
type Attribute struct {
	ID        uint   `json:"attribute_id" gorm:"primaryKey"`
	ContentID uint   `json:"content_id" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`
	Value     string `json:"value" gorm:"type:jsonb"`
	CreatedAt time.Time
}

// Child table model | Children relationship table (self-referencing Content)
type Child struct {
	ParentID uint
	ChildID  uint
}
