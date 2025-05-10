package postgres

import "gorm.io/gorm"

type PostgresInterface interface {
	CreatePost(userId int, text string, image []byte) (int, error)
	DeletePost(userId int, postId int) error
	CreateComment(userId, postId int, text string) error
	DeleteComment(userId int, commentId int) error
	GetPost(postId int) (*Post, error)
	GetPosts(userIds []int, page int) ([]Post, error)
	GetCountCommentsAndLikes(postIds []int) (map[int]int, map[int]int, error)
	GetComments(postId, page int) ([]Comment, error)
	Like(postId int, userId int) error
}

var ProductionPostgresInterface PostgresInterface = &RealPostgres{instance: initPostgres()}

type RealPostgres struct {
	instance *gorm.DB
}
