package postgres

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id             int        `gorm:"primaryKey;autoIncrement"`
	Login          string     `gorm:"not null;size:30"`
	Name           string     `gorm:"not null;size:30"`
	Email          string     `gorm:"not null;size:30;unique"`
	ConfirmedEmail bool       `gorm:"default:false"`
	Password       string     `gorm:"not null;size:30"`
	Birthday       *time.Time `gorm:"default:null"`
	Phone          string     `gorm:"default:null"`
	City           string     `gorm:"default:null"`
	Image          []byte     `gorm:"type:bytea"` // Хранение изображения в виде бинарных данных
}

func CompareUsers(user1, user2 User) bool {
	return user1.Id == user2.Id
}

type Friend struct {
	Id       int  `gorm:"primaryKey;autoIncrement"`
	User1Id  int  `gorm:"not null;index"`
	User2Id  int  `gorm:"not null;index"`
	Accepted bool `gorm:"not null"`

	User1 *User `gorm:"foreignKey:User1Id;references:Id"`
	User2 *User `gorm:"foreignKey:User2Id;references:Id"`
}
