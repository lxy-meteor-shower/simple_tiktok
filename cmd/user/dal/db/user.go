package db

import (
	"context"
	"demo/tiktok/pkg/constants"
	"strings"

	"gorm.io/gorm"
)

type User struct {
	Id              int64  `json:"id"`
	Name            string `json:"name"`
	Password        string `json:"password"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

func (u *User) TableName() string {
	return constants.UserTableName
}

func CreateUser(ctx context.Context, user *User) error {
	return DB.WithContext(ctx).Create(user).Error
}

func QueryUser(ctx context.Context, username string) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("name = ?", username).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func MGetUser(ctx context.Context, userid int64) ([]*User, error) {
	res := make([]*User, 0)
	if err := DB.WithContext(ctx).Where("id = ?", userid).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

func UpdateUser(ctx context.Context, userid int64, column string, value int) error {
	var builder strings.Builder
	builder.WriteString(column)
	builder.WriteString(" + ?")
	return DB.WithContext(ctx).Model(&User{}).Where("id = ?", userid).UpdateColumn(column, gorm.Expr(builder.String(), 1)).Error
}
