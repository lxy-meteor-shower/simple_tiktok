package service

import (
	"context"
	"demo/tiktok/cmd/user/dal/db"
	"demo/tiktok/kitex_gen/userdemo"
)

type IncUserVideoService struct {
	ctx context.Context
}

func NewIncUserVideoService(ctx context.Context) *IncUserVideoService {
	return &IncUserVideoService{ctx: ctx}
}

func (s *IncUserVideoService) IncUserVideo(req *userdemo.ChangeUserRequest) error {
	return db.UpdateUser(s.ctx, req.UserId, "work_count", 1)
}
