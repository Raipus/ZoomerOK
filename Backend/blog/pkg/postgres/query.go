package postgres

import (
	"errors"
)

func (r *PostgresRepository) CreatePost(userId int, text string, photo []byte) error {
	var post Post
	post.UserId = userId
	post.Text = text
	post.Photo = photo
	return r.db.Create(post).Error
}

func (r *PostgresRepository) DeletePost(userId int, postId int) error {
	var post Post
	if err := r.db.First(&post, postId).Error; err != nil {
		return err
	}
	if post.UserId != userId {
		return errors.New("unauthorized to delete this post")
	}
	return r.db.Delete(&post).Error
}

func (r *PostgresRepository) CreateComment(userId int, text string) error {
	var comment Comment
	comment.UserId = userId
	comment.Text = text
	return r.db.Create(comment).Error
}

func (r *PostgresRepository) DeleteComment(userId int, commentId int) error {
	var comment Comment
	if err := r.db.First(&comment, commentId).Error; err != nil {
		return err
	}
	if comment.UserId != userId {
		return errors.New("unauthorized to delete this comment")
	}
	return r.db.Delete(&comment).Error
}

func (r *PostgresRepository) GetPost(postId int) (*Post, error) {
	var post Post
	if err := r.db.First(&post, postId).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostgresRepository) GetPosts(userIds []int) ([]Post, error) {
	var posts []Post
	if err := r.db.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostgresRepository) Like(userId int, postId int) error {
	like := Like{PostId: postId, UserId: userId}
	return r.db.Create(&like).Error
}
