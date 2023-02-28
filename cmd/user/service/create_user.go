package service

import (
	"context"
	"crypto/md5"
	"demo/tiktok/cmd/user/dal/db"
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/pkg/errno"
	"fmt"
	"io"
)

type CreateUserService struct {
	ctx context.Context
}

func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

func (s *CreateUserService) CreateUser(req *userdemo.CreateUserRequest) error {
	users, err := db.QueryUser(s.ctx, req.UserName)
	if err != nil {
		return err
	}
	if len(users) != 0 {
		return errno.UserAlreadyExistErr
	}

	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))
	return db.CreateUser(s.ctx, &db.User{
		Name:     req.UserName,
		Password: passWord,
	})
}
