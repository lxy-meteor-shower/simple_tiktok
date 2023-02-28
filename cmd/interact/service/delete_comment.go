package service

import (
	"context"
	"demo/tiktok/cmd/interact/dal/db"
	"demo/tiktok/kitex_gen/interactdemo"
)

type DeleteCommentService struct {
	ctx context.Context
}

func NewDeleteCommentService(ctx context.Context) *DeleteCommentService {
	return &DeleteCommentService{ctx: ctx}
}

func (s *DeleteCommentService) DeleteComment(req *interactdemo.DeleteCommentRequest) error {
	return db.DeleteComment(s.ctx, req.CommentId)
}
