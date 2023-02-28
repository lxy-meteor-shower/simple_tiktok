package rpc

import (
	"context"
	"demo/tiktok/kitex_gen/interactdemo"
	"demo/tiktok/kitex_gen/interactdemo/interactservice"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var interactClient interactservice.Client

func InitInteractRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := interactservice.NewClient(
		constants.InteractServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	interactClient = c
}

func IsFavorite(ctx context.Context, req *interactdemo.CheckFavoriteRequest) (bool, error) {
	resp, err := interactClient.CheckFavorite(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return false, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.IsFavorite, nil
}
