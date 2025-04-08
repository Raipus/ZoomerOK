package postgres

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id             int        `gorm:"type:id;primaryKey;not null"`
	Name           string     `gorm:"not null;size:30"`
	Email          string     `gorm:"not null;size:30;unique"`
	ConfirmedEmail bool       `gorm:"default:false"`
	Password       string     `gorm:"not null;size:30"`
	Birthday       *time.Time `gorm:"default:null"`
	Phone          string     `gorm:"default:null"`
	City           string     `gorm:"default:null"`
	Image          []byte     `gorm:"type:bytea"` // Хранение изображения в виде бинарных данных
}

type Friend struct {
	Id       string `gorm:"type:id;primaryKey"`
	User1Id  int    `gorm:"not null;index;type:id"`
	User2Id  int    `gorm:"not null;index;type:id"`
	Accepted bool   `gorm:"not null"`

	User1 *User `gorm:"foreignKey:User1Id;references:Id"`
	User2 *User `gorm:"foreignKey:User2Id;references:Id"`
}
