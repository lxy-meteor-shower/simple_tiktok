package service

import (
	"context"
	"demo/tiktok/cmd/video/dal/db"
	"demo/tiktok/kitex_gen/videodemo"
)

type DescVideoFavoriteService struct {
	ctx context.Context
}

func NewDescVideoFavoriteService(ctx context.Context) *DescVideoFavoriteService {
	return &DescVideoFavoriteService{ctx: ctx}
}

func (s *DescVideoFavoriteService) DescVideoFavorite(req *videodemo.ChangeVideoRequest) error {
	return db.UpdateVideo(s.ctx, req.VideoId, "favorite_count", -1)
}
