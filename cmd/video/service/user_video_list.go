package service

import (
	"context"
	"demo/tiktok/cmd/video/dal/db"
	"demo/tiktok/cmd/video/pack"
	"demo/tiktok/kitex_gen/videodemo"
)

type UserVideoListService struct {
	ctx context.Context
}

func NewUserVideoListService(ctx context.Context) *UserVideoListService {
	return &UserVideoListService{ctx: ctx}
}

func (s *UserVideoListService) UserVideoList(req *videodemo.GetUserVideoRequest) ([]*videodemo.Video, error) {
	videos, err := db.GetUserVideo(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.VideoList(s.ctx, videos, false, req.UserId)
}
