package service

import (
	"context"
	"demo/tiktok/cmd/interact/dal/db"
	"demo/tiktok/kitex_gen/interactdemo"
)

type CreateFavoriteService struct {
	ctx context.Context
}

func NewCreateFavoriteService(ctx context.Context) *CreateFavoriteService {
	return &CreateFavoriteService{ctx: ctx}
}

func (s *CreateFavoriteService) CreateFavorite(req *interactdemo.FavoriteRequest) error {
	return db.CreateFavorite(s.ctx, &db.Favorite{
		UserID:  req.UserId,
		VideoID: req.VideoId,
	})
}
