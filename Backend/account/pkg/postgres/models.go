package postgres

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID     string     `gorm:"type:uuid;default:uuid_generate_v4();unique;not null"`
	Name     string     `gorm:"not null;size:30"`
	Email    string     `gorm:"not null;size:30;unique"`
	Password string     `gorm:"not null;size:30"`
	Birthday *time.Time `gorm:"default:null"`
	Phone    string     `gorm:"default:null"`
	City     string     `gorm:"default:null"`
	Image    []byte     `gorm:"type:bytea"` // Хранение изображения в виде бинарных данных
}

type Friend struct {
	UUID      string `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	User1UUID string `gorm:"not null;index;type:uuid"`
	User2UUID string `gorm:"not null;index;type:uuid"`
	Accepted  bool   `gorm:"not null"`

	User1 *User `gorm:"foreignKey:User1UUID;references:UUID"`
	User2 *User `gorm:"foreignKey:User2UUID;references:UUID"`
}
