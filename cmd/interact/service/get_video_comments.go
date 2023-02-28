package service

import (
	"context"
	"demo/tiktok/cmd/interact/dal/db"
	"demo/tiktok/cmd/interact/pack"
	"demo/tiktok/kitex_gen/interactdemo"
)

type GetVideoCommentService struct {
	ctx context.Context
}

func NewGetVideoCommentService(ctx context.Context) *GetVideoCommentService {
	return &GetVideoCommentService{ctx: ctx}
}

func (s *GetVideoCommentService) GetVideoComment(req *interactdemo.GetVideoCommentRequest) ([]*interactdemo.Comment, error) {
	comments, err := db.GetVideoComment(s.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return pack.CommentList(s.ctx, comments)
}
