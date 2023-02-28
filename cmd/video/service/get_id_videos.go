package service

import (
	"context"
	"demo/tiktok/cmd/video/dal/db"
	"demo/tiktok/cmd/video/pack"
	"demo/tiktok/kitex_gen/videodemo"
)

type GetIdVideosService struct {
	ctx context.Context
}

func NewGetIdVideosService(ctx context.Context) *GetIdVideosService {
	return &GetIdVideosService{ctx: ctx}
}

func (s *GetIdVideosService) GetIdVideos(req *videodemo.GetFavoriteVideoRequest) ([]*videodemo.Video, error) {
	videos, err := db.GetIdVideo(s.ctx, req.VideoIds)
	if err != nil {
		return nil, err
	}
	return pack.VideoList(s.ctx, videos, true, 0)
}
