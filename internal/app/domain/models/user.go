package models

import (
	"gorm.io/gorm"
)

// User represents a user in the system.
type User struct {
	gorm.Model
	ID       string `gorm:"column:id;primaryKey"`
	Name     string `gorm:"column:name"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Email    string `gorm:"column:email"`
	Phone    string `gorm:"column:phone"`
}
