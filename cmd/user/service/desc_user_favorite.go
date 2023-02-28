package service

import (
	"context"
	"demo/tiktok/cmd/user/dal/db"
	"demo/tiktok/kitex_gen/userdemo"
)

type DescUserFavoriteService struct {
	ctx context.Context
}

func NewDescUserFavoriteService(ctx context.Context) *DescUserFavoriteService {
	return &DescUserFavoriteService{ctx: ctx}
}

func (s *DescUserFavoriteService) DescUserFavorite(req *userdemo.ChangeUserRequest) error {
	return db.UpdateUser(s.ctx, req.UserId, "favorite_count", -1)
}
