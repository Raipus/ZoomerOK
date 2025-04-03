package postgres

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string     `gorm:"type:uuid;default:uuid_generate_v4();unique;not null"`
	Email    string     `gorm:"not null;unique"`
	Password string     `gorm:"not null"`
	Name     string     `gorm:"not null"`
	Birthday *time.Time `gorm:"not null"`
	Phone    string     `gorm:"not null"`
	City     string     `gorm:"not null"`
}

type Friend struct {
	UUID      string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	User1UUID string `gorm:"not null;index;type:uuid"`
	User2UUID string `gorm:"not null;index;type:uuid"`
	Accepted  bool   `gorm:"not null"`

	User1 *User `gorm:"foreignKey:User1UUID;references:UUID"`
	User2 *User `gorm:"foreignKey:User2UUID;references:UUID"`
}
