package rpc

import (
	"context"
	"demo/tiktok/kitex_gen/videodemo"
	"demo/tiktok/kitex_gen/videodemo/videoservice"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"

	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var videoClient videoservice.Client

func InitVideoRPC() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := videoservice.NewClient(
		constants.VideoServiceName,
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	videoClient = c
}

func GetIdVideos(ctx context.Context, req *videodemo.GetFavoriteVideoRequest) ([]*videodemo.Video, error) {
	resp, err := videoClient.GetFavoriteVideo(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}
