package types

import (
	"time"
)

// Type table model
type Type struct {
	ID   uint   `json:"type_id" gorm:"primaryKey;autoincrement"`
	Name string `json:"name" gorm:"not null; unique"`
}

// Content table model
type Content struct {
	ID         uint `json:"content_id" gorm:"primaryKey;autoincrement"`
	TypeID     uint `json:"type_id" gorm:"not null"`
	Type       Type `gorm:"foreignKey:TypeID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Data       Data        `json:"data" gorm:"foreignKey:ContentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attributes []Attribute `json:"attributes" gorm:"foreignKey:ContentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Children   []*Content  `json:"children" gorm:"many2many:content_children;joinForeignKey:ParentID;joinReferences:ChildID"`
}

// Data table model
type Data struct {
	ID        uint     `json:"data_id" gorm:"primaryKey;autoincrement"`
	ContentID uint     `json:"content_id" gorm:"not null"`
	Text      []string `json:"text" gorm:"type:text[]"`
	ImageUrls []string `json:"image" gorm:"type:text[]"`
	PdfUrls   []string `json:"pdf" gorm:"type:text[]"`
	Code      []string `json:"code" gorm:"type:text[]"`
	CreatedAt time.Time
	Content   *Content `gorm:"foreignKey:ContentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Foreign key back to content
}

// Attribute table model
type Attribute struct {
	ID        uint   `json:"attribute_id" gorm:"primaryKey;autoincrement"`
	ContentID uint   `json:"content_id" gorm:"not null"`
	Name      string `json:"name" gorm:"not null"`
	Value     string `json:"value" gorm:"type:jsonb"`
	CreatedAt time.Time
	Content   *Content `gorm:"foreignKey:ContentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Foreign key back to content
}

// Child table model | Children relationship table (self-referencing Content)
type Child struct {
	ParentID uint     `gorm:"primaryKey"`
	ChildID  uint     `gorm:"primaryKey"`
	Parent   *Content `gorm:"foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Child    *Content `gorm:"foreignKey:ChildID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
