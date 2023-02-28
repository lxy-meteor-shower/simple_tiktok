package service

import (
	"context"
	"demo/tiktok/cmd/video/dal/db"
	"demo/tiktok/kitex_gen/videodemo"
)

type CreateVideoService struct {
	ctx context.Context
}

func NewCreateVideoService(ctx context.Context) *CreateVideoService {
	return &CreateVideoService{ctx: ctx}
}

func (s *CreateVideoService) CreateVideo(req *videodemo.CreateVideoRequest) error {
	return db.CreateVideo(s.ctx, &db.Video{
		Author:   req.UserId,
		PlayURL:  req.PlayUrl,
		CoverURL: req.CoverUrl,
		Title:    req.Title,
	})
}
