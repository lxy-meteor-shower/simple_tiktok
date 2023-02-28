package service

import (
	"context"
	"demo/tiktok/cmd/user/dal/db"
	"demo/tiktok/kitex_gen/userdemo"
)

type DescUserFavoritedService struct {
	ctx context.Context
}

func NewDescUserFavoritedService(ctx context.Context) *DescUserFavoritedService {
	return &DescUserFavoritedService{ctx: ctx}
}

func (s *DescUserFavoritedService) DescUserFavorited(req *userdemo.ChangeUserRequest) error {
	return db.UpdateUser(s.ctx, req.UserId, "total_favorited", -1)
}
