package db

import (
	"context"
	"demo/tiktok/pkg/constants"
	"time"
)

type Favorite struct {
	UserID   int64     `json:"user_id" gorm:"primary_key"`
	VideoID  int64     `json:"video_id" gorm:"primary_key"`
	CreateAt time.Time `json:"create_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (v *Favorite) TableName() string {
	return constants.FavoriteTableName
}

func CreateFavorite(ctx context.Context, favorite *Favorite) error {
	return DB.WithContext(ctx).Create(favorite).Error
}

func DeleteFavorite(ctx context.Context, uid int64, vid int64) error {
	return DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", uid, vid).Delete(&Favorite{}).Error
}

func GetFavoriteVideoIds(ctx context.Context, uid int64) ([]int64, error) {
	res := make([]int64, 0)
	if err := DB.WithContext(ctx).Model(&Favorite{}).Where("user_id = ?", uid).Select("video_id").Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func GetOneFavorite(ctx context.Context, uid int64, vid int64) ([]*Favorite, error) {
	res := make([]*Favorite, 0)
	if err := DB.WithContext(ctx).Where("user_id = ? AND video_id = ?", uid, vid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}
