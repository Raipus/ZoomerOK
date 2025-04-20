package postgres

import (
	"errors"
)

func (Instance *RealPostgres) CreatePost(userId int, text string, photo []byte) error {
	var post Post
	post.UserId = userId
	post.Text = text
	post.Photo = photo
	return Instance.instance.Create(post).Error
}

func (Instance *RealPostgres) DeletePost(userId int, postId int) error {
	var post Post
	if err := Instance.instance.First(&post, postId).Error; err != nil {
		return err
	}
	if post.UserId != userId {
		return errors.New("unauthorized to delete this post")
	}
	return Instance.instance.Delete(&post).Error
}

func (Instance *RealPostgres) CreateComment(userId int, text string) error {
	var comment Comment
	comment.UserId = userId
	comment.Text = text
	return Instance.instance.Create(comment).Error
}

func (Instance *RealPostgres) DeleteComment(userId int, commentId int) error {
	var comment Comment
	if err := Instance.instance.First(&comment, commentId).Error; err != nil {
		return err
	}
	if comment.UserId != userId {
		return errors.New("unauthorized to delete this comment")
	}
	return Instance.instance.Delete(&comment).Error
}

func (Instance *RealPostgres) GetPost(postId int) (*Post, error) {
	var post Post
	if err := Instance.instance.First(&post, postId).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (Instance *RealPostgres) GetPosts(userId int) ([]Post, error) {
	var posts []Post
	if err := Instance.instance.Where("user_id = ?", userId).Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (Instance *RealPostgres) Like(userId int, postId int) error {
	like := Like{PostId: postId, UserId: userId}
	return Instance.instance.Create(&like).Error
}
