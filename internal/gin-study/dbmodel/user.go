// Package dbmodel Code generated by sql2struct. https://github.com/starfishs/sql2struct
package dbmodel

import "time"

type User struct {
	ID        int64     `gorm:"column:id" json:"id,omitempty"`
	Username  string    `gorm:"column:username" json:"username,omitempty"`
	Password  string    `gorm:"column:password" json:"password,omitempty"`
	Email     string    `gorm:"column:email" json:"email,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

// TableName the name of table in database
func (t *User) TableName() string {
	return "user"
}