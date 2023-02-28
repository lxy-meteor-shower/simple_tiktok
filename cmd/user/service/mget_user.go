package service

import (
	"context"
	"demo/tiktok/cmd/user/dal/db"
	"demo/tiktok/cmd/user/pack"
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/pkg/errno"
)

type MGetUserService struct {
	ctx context.Context
}

func NewMGetUserService(ctx context.Context) *MGetUserService {
	return &MGetUserService{ctx: ctx}
}

func (s *MGetUserService) MGetUser(req *userdemo.MGetUserRequest) (*userdemo.User, error) {
	users, err := db.MGetUser(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, errno.ParamErr
	}
	return pack.User(users[0]), nil
}
