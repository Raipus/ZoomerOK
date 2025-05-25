package postgres

import (
	"errors"
	"time"

	"github.com/Raipus/ZoomerOK/blog/pkg/config"
)

func (Instance *RealPostgres) CreatePost(userId int, text string, image []byte) (int, error) {
	var post Post
	post.UserId = userId
	post.Text = text
	post.Image = image
	now := time.Now()
	post.Time = &now

	result := Instance.instance.Create(&post)
	if result.Error != nil {
		return 0, result.Error
	}
	return post.Id, nil
}

func (Instance *RealPostgres) DeletePost(userId int, postId int) error {
	var post Post
	if err := Instance.instance.First(&post, postId).Error; err != nil {
		return errors.New("пользователь не имеет права на данную операцию")
	}
	if post.UserId != userId {
		return errors.New("пользователь не имеет права на данную операцию")
	}
	return Instance.instance.Delete(&post).Error
}

func (Instance *RealPostgres) CreateComment(userId, postId int, text string) error {
	var comment Comment
	comment.UserId = userId
	comment.PostId = postId
	comment.Text = text
	now := time.Now()
	comment.Time = &now
	return Instance.instance.Create(&comment).Error
}

func (Instance *RealPostgres) DeleteComment(userId int, commentId int) error {
	var comment Comment
	if err := Instance.instance.First(&comment, commentId).Error; err != nil {
		return errors.New("пользователь не имеет права на данную операцию")
	}
	if comment.UserId != userId {
		return errors.New("пользователь не имеет права на данную операцию")
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

func (Instance *RealPostgres) GetPosts(userIds []int, page int) ([]Post, error) {
	var posts []Post
	offset := (page - 1) * config.Config.PageSize
	if err := Instance.instance.Where("user_id IN ?", userIds).
		Order("created_at DESC").
		Limit(config.Config.PageSize).
		Offset(offset).
		Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (Instance *RealPostgres) GetCountCommentsAndLikes(postIds []int) (map[int]int, map[int]int, error) {
	commentCountMap := make(map[int]int)
	likeCountMap := make(map[int]int)

	var commentCounts []struct {
		PostId       int
		CommentCount int64
	}

	err := Instance.instance.Model(&Comment{}).
		Select("post_id, count(*) as comment_count").
		Where("post_id IN ?", postIds).
		Group("post_id").
		Scan(&commentCounts).Error
	if err != nil {
		return nil, nil, err
	}

	for _, cc := range commentCounts {
		commentCountMap[cc.PostId] = int(cc.CommentCount)
	}

	var likeCounts []struct {
		PostId    int
		LikeCount int64
	}

	err = Instance.instance.Model(&Like{}).
		Select("post_id, count(*) as like_count").
		Where("post_id IN ?", postIds).
		Group("post_id").
		Scan(&likeCounts).Error

	if err != nil {
		return nil, nil, err
	}

	for _, lc := range likeCounts {
		likeCountMap[lc.PostId] = int(lc.LikeCount)
	}

	return commentCountMap, likeCountMap, nil
}

func (Instance *RealPostgres) GetComments(postId, page int) ([]Comment, error) {
	var comments []Comment
	offset := (page - 1) * config.Config.PageSize
	if err := Instance.instance.Where("post_id = ?", postId).
		Order("created_at DESC").
		Limit(config.Config.PageSize).
		Offset(offset).
		Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (Instance *RealPostgres) Like(userId int, postId int) error {
	var like Like
	if err := Instance.instance.Where("user_id = ? AND post_id = ?", userId, postId).First(&like).Error; err != nil {
		newLike := Like{PostId: postId, UserId: userId}
		return Instance.instance.Create(&newLike).Error
	} else {
		return Instance.instance.Delete(&like).Error
	}
}
