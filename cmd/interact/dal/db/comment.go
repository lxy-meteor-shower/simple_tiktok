package db

import (
	"context"
	"demo/tiktok/pkg/constants"
	"time"
)

type Comment struct {
	Id       int64     `json:"id"`
	UserID   int64     `json:"user_id"`
	VideoID  int64     `json:"video_id"`
	Content  string    `json:"content"`
	CreateAt time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (c *Comment) TableName() string {
	return constants.CommentTableName
}

func CreateComment(ctx context.Context, comment *Comment) (*Comment, error) {
	var c Comment
	if err := DB.WithContext(ctx).Create(comment).Scan(&c).Error; err != nil {
		return nil, err
	}
	return &c, nil
}

func DeleteComment(ctx context.Context, id int64) error {
	return DB.WithContext(ctx).Delete(&Comment{}, id).Error
}

func GetVideoComment(ctx context.Context, vid int64) ([]*Comment, error) {
	res := make([]*Comment, 0)
	if err := DB.WithContext(ctx).Where("video_id = ?", vid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
