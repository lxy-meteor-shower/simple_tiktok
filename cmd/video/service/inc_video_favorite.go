package service

import (
	"context"
	"demo/tiktok/cmd/video/dal/db"
	"demo/tiktok/kitex_gen/videodemo"
)

type IncVideoFavoriteService struct {
	ctx context.Context
}

func NewIncVideoFavoriteService(ctx context.Context) *IncVideoFavoriteService {
	return &IncVideoFavoriteService{ctx: ctx}
}

func (s *IncVideoFavoriteService) IncVideoFavorite(req *videodemo.ChangeVideoRequest) error {
	return db.UpdateVideo(s.ctx, req.VideoId, "favorite_count", 1)
}
