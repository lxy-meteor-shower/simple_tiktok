package rpc

import (
	"context"
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/kitex_gen/userdemo/userservice"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		constants.UserServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func CreateUser(ctx context.Context, req *userdemo.CreateUserRequest) (int64, error) {
	resp, err := userClient.CreateUser(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}

func CheckUser(ctx context.Context, req *userdemo.CheckUserRequest) (int64, error) {
	resp, err := userClient.CheckUser(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserId, nil
}

func GetUserInfo(ctx context.Context, req *userdemo.MGetUserRequest) (*userdemo.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.UserInfo, nil
}

func IncFavorited(ctx context.Context, req *userdemo.ChangeUserRequest) error {
	resp, err := userClient.IncFavorited(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func IncVideo(ctx context.Context, req *userdemo.ChangeUserRequest) error {
	resp, err := userClient.IncVideo(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func IncFavorite(ctx context.Context, req *userdemo.ChangeUserRequest) error {
	resp, err := userClient.IncFavorite(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func DescFavorited(ctx context.Context, req *userdemo.ChangeUserRequest) error {
	resp, err := userClient.DescFavorited(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func DescFavorite(ctx context.Context, req *userdemo.ChangeUserRequest) error {
	resp, err := userClient.DescFavorite(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}
