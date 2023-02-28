package service

import (
	"context"
	"demo/tiktok/cmd/user/dal/db"
	"demo/tiktok/kitex_gen/userdemo"
)

type IncUserFavoriteService struct {
	ctx context.Context
}

func NewIncUserFavoriteService(ctx context.Context) *IncUserFavoriteService {
	return &IncUserFavoriteService{ctx: ctx}
}

func (s *IncUserFavoriteService) IncUserFavorite(req *userdemo.ChangeUserRequest) error {
	return db.UpdateUser(s.ctx, req.UserId, "favorite_count", 1)
}
