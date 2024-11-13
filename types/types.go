package types

import (
	"github.com/lib/pq"
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
	TypeID     uint `json:"type_id" gorm:"foreignKey:Type;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DataID     uint        `json:"data" gorm:"foreignKey:ContentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Attributes []Attribute `json:"attributes" gorm:"foreignKey:ContentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Children   []Content   `json:"children" gorm:"many2many:content_children;joinForeignKey:ParentID;joinReferences:ChildID"`
}

// Data table model
type Data struct {
	ID        uint           `json:"data_id" gorm:"primaryKey;autoincrement"`
	ContentID uint           `json:"content_id" gorm:"foreignKey:ContentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Text      pq.StringArray `json:"text" gorm:"type:text[]"`
	ImageUrls pq.StringArray `json:"image" gorm:"type:text[]"`
	PdfUrls   pq.StringArray `json:"pdf" gorm:"type:text[]"`
	Code      pq.StringArray `json:"code" gorm:"type:text[]"`
	CreatedAt time.Time
}

// Attribute table model
type Attribute struct {
	ID        uint    `json:"attribute_id" gorm:"primaryKey;autoincrement"`
	ContentID uint    `json:"content_id" gorm:"foreignKey:ContentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;not null"`
	Name      string  `json:"name" gorm:"not null"`
	Value     float32 `json:"value" gorm:"type:jsonb"`
	CreatedAt time.Time
}

// Child table model | Children relationship table (self-referencing Content)
type Child struct {
	ParentID uint `gorm:"primaryKey;foreignKey:ParentID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	ChildID  uint `gorm:"primaryKey;foreignKey:ChildID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
