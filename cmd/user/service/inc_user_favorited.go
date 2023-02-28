package service

import (
	"context"
	"demo/tiktok/cmd/user/dal/db"
	"demo/tiktok/kitex_gen/userdemo"
)

type IncUserFavoritedService struct {
	ctx context.Context
}

func NewIncUserFavoritedService(ctx context.Context) *IncUserFavoritedService {
	return &IncUserFavoritedService{ctx: ctx}
}

func (s *IncUserFavoritedService) IncUserFavorited(req *userdemo.ChangeUserRequest) error {
	return db.UpdateUser(s.ctx, req.UserId, "total_favorited", 1)
}
