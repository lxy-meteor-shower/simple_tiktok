package service

import (
	"context"
	"demo/tiktok/cmd/video/dal/db"
	"demo/tiktok/kitex_gen/videodemo"
)

type DescVideoCommentService struct {
	ctx context.Context
}

func NewDescVideoCommentService(ctx context.Context) *DescVideoCommentService {
	return &DescVideoCommentService{ctx: ctx}
}

func (s *DescVideoCommentService) DescVideoComment(req *videodemo.ChangeVideoRequest) error {
	return db.UpdateVideo(s.ctx, req.VideoId, "comment_count", -1)
}
