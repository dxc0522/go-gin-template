// Package model Code generated by sql2struct. https://github.com/starfishs/sql2struct
package model

import "time"

type User struct {
	ID        int64     `gorm:"column:id" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Email     string    `gorm:"column:email" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName the name of table in database
func (t *User) TableName() string {
	return "user"
}