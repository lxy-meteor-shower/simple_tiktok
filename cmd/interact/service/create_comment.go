package service

import (
	"context"
	"demo/tiktok/cmd/interact/dal/db"
	"demo/tiktok/cmd/interact/pack"
	"demo/tiktok/kitex_gen/interactdemo"
)

type CreateCommentService struct {
	ctx context.Context
}

func NewCreateCommentService(ctx context.Context) *CreateCommentService {
	return &CreateCommentService{ctx: ctx}
}

func (s *CreateCommentService) CreateComment(req *interactdemo.CreateCommentRequest) (*interactdemo.Comment, error) {
	comment, err := db.CreateComment(s.ctx, &db.Comment{
		UserID:  req.UserId,
		VideoID: req.VideoId,
		Content: req.CommentText,
	})
	if err != nil {
		return nil, err
	}
	return pack.Comment(s.ctx, comment)
}
