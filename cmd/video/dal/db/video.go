package db

import (
	"context"
	"demo/tiktok/pkg/constants"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Video struct {
	Id            int64     `json:"id"`
	Author        int64     `json:"author"`
	PlayURL       string    `json:"play_url"`
	CoverURL      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	Title         string    `json:"title"`
	CreateAt      time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (v *Video) TableName() string {
	return constants.VideoTableName
}

func CreateVideo(ctx context.Context, video *Video) error {
	return DB.WithContext(ctx).Create(video).Error
}

func GetUserVideo(ctx context.Context, userid int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("author = ?", userid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetFeed(ctx context.Context, last_time string) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("create_at < ?", last_time).Order("create_at desc").Limit(30).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetIdVideo(ctx context.Context, ids []int64) ([]*Video, error) {
	res := make([]*Video, 0)
	if err := DB.WithContext(ctx).Where("id in ?", ids).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateVideo(ctx context.Context, video_id int64, column string, value int) error {
	var builder strings.Builder
	builder.WriteString(column)
	builder.WriteString(" + ?")
	return DB.WithContext(ctx).Model(&Video{}).Where("id = ?", video_id).UpdateColumn(column, gorm.Expr(builder.String(), value)).Error
}
