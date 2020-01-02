package gormodel

import (
	"time"
)

// Model likes gorm.Model but with json little camel-case keys.
type Model struct {
	ID        uint       `gorm:"PRIMARY_KEY" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
