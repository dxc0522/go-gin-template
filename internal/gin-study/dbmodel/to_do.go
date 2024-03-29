// Package dbmodel Code generated by sql2struct. https://github.com/starfishs/sql2struct
package dbmodel

import "time"

type ToDo struct {
	ID        int64     `gorm:"column:id" json:"id,omitempty"`
	Label     string    `gorm:"column:label" json:"label,omitempty"`
	Value     string    `gorm:"column:value" json:"value,omitempty"`
	Disable   uint8     `gorm:"column:disable" json:"disable,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

// TableName the name of table in database
func (t *ToDo) TableName() string {
	return "to_do"
}
