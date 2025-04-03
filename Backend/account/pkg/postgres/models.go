package postgres

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string `gorm:"primaryKey"`
	Password string
	Name     string
	Birthday *time.Time
	Phone    string
	City     string
}

type Friend struct {
	UUID     string `gorm:"primaryKey"`
	User1    User
	User2    User
	Accepted bool
}
