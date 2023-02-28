package service

import (
	"context"
	"demo/tiktok/cmd/video/dal/db"
	"demo/tiktok/kitex_gen/videodemo"
)

type IncVideoCommentService struct {
	ctx context.Context
}

func NewIncVideoCommentService(ctx context.Context) *IncVideoCommentService {
	return &IncVideoCommentService{ctx: ctx}
}

func (s *IncVideoCommentService) IncVideoComment(req *videodemo.ChangeVideoRequest) error {
	return db.UpdateVideo(s.ctx, req.VideoId, "comment_count", 1)
}
