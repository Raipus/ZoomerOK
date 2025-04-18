package postgres

import "gorm.io/gorm"

type PostgresInterface interface {
	CreatePost(userId int, post *Post) error
	DeletePost(userId int, postId int) error
	CreateComment(userId int, comment *Comment) error
	DeleteComment(userId int, commentId int) error
	GetPost(postId int) (*Post, error)
	GetPosts(userId int) ([]Post, error)
	Like(postId int, userId int) error
}

var ProductionPostgresInterface PostgresInterface = &RealPostgres{instance: initPostgres()}

type RealPostgres struct {
	instance *gorm.DB
}
