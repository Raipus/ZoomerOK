package postgres

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Id     int `gorm:"type:id;primaryKey;not null"`
	UserId int
	Text   string     `gorm:"not null"`
	Image  []byte     `gorm:"type:bytea"`
	Time   *time.Time `gorm:"not null"`
}

type Comment struct {
	gorm.Model
	Id     int        `gorm:"primaryKey;not null"`
	UserId int        `gorm:"not null"`
	PostId int        `gorm:"not null"`
	Text   string     `gorm:"not null"`
	Time   *time.Time `gorm:"not null"`
}

type Like struct {
	gorm.Model
	Id     int `gorm:"primaryKey;not null"`
	UserId int
	PostId int `gorm:"not null"`
}
