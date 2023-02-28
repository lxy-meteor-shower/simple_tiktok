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

func InitUserRPC() {
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
