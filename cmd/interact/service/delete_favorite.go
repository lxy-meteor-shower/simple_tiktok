package service

import (
	"context"
	"demo/tiktok/cmd/interact/dal/db"
	"demo/tiktok/kitex_gen/interactdemo"
)

type DeleteFavoriteService struct {
	ctx context.Context
}

func NewDeleteFavoriteService(ctx context.Context) *DeleteFavoriteService {
	return &DeleteFavoriteService{ctx: ctx}
}

func (s *DeleteFavoriteService) DeleteFavorite(req *interactdemo.FavoriteRequest) error {
	return db.DeleteFavorite(s.ctx, req.UserId, req.VideoId)
}
