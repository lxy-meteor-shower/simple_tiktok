package main

import (
	"context"
	"demo/tiktok/cmd/user/pack"
	"demo/tiktok/cmd/user/service"
	userdemo "demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (resp *userdemo.CreateUserResponse, err error) {
	resp = new(userdemo.CreateUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// CheckUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (resp *userdemo.CheckUserResponse, err error) {
	resp = new(userdemo.CheckUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		resp.BaseResp = pack.BuildBaseResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewCheckUserService(ctx).CheckUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// MGetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) MGetUser(ctx context.Context, req *userdemo.MGetUserRequest) (resp *userdemo.MGetUserResponse, err error) {
	resp = new(userdemo.MGetUserResponse)

	user, err := service.NewMGetUserService(ctx).MGetUser(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.UserInfo = user
	return resp, nil
}

// IncFavorited implements the UserServiceImpl interface.
func (s *UserServiceImpl) IncFavorited(ctx context.Context, req *userdemo.ChangeUserRequest) (resp *userdemo.ChangeUserResponse, err error) {
	resp = new(userdemo.ChangeUserResponse)

	err = service.NewIncUserFavoritedService(ctx).IncUserFavorited(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// IncVideo implements the UserServiceImpl interface.
func (s *UserServiceImpl) IncVideo(ctx context.Context, req *userdemo.ChangeUserRequest) (resp *userdemo.ChangeUserResponse, err error) {
	resp = new(userdemo.ChangeUserResponse)

	err = service.NewIncUserVideoService(ctx).IncUserVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// IncFavorite implements the UserServiceImpl interface.
func (s *UserServiceImpl) IncFavorite(ctx context.Context, req *userdemo.ChangeUserRequest) (resp *userdemo.ChangeUserResponse, err error) {
	resp = new(userdemo.ChangeUserResponse)

	err = service.NewIncUserFavoriteService(ctx).IncUserFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DescFavorited implements the UserServiceImpl interface.
func (s *UserServiceImpl) DescFavorited(ctx context.Context, req *userdemo.ChangeUserRequest) (resp *userdemo.ChangeUserResponse, err error) {
	resp = new(userdemo.ChangeUserResponse)

	err = service.NewDescUserFavoritedService(ctx).DescUserFavorited(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DescFavorite implements the UserServiceImpl interface.
func (s *UserServiceImpl) DescFavorite(ctx context.Context, req *userdemo.ChangeUserRequest) (resp *userdemo.ChangeUserResponse, err error) {
	resp = new(userdemo.ChangeUserResponse)

	err = service.NewDescUserFavoriteService(ctx).DescUserFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}

	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}
