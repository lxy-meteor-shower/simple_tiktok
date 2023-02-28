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

func initVideoRPC() {
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

func CreateVideo(ctx context.Context, req *videodemo.CreateVideoRequest) error {
	resp, err := videoClient.CreateVideo(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetUserVideoList(ctx context.Context, req *videodemo.GetUserVideoRequest) ([]*videodemo.Video, error) {
	resp, err := videoClient.GetUserVideo(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}

func GetFeed(ctx context.Context, req *videodemo.GetFeedRequest) ([]*videodemo.Video, string, error) {
	resp, err := videoClient.GetFeed(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, "", errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, resp.NextTime, nil
}

func IncVideoFavorite(ctx context.Context, req *videodemo.ChangeVideoRequest) error {
	resp, err := videoClient.IncFavoriteCount(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func IncVideoComment(ctx context.Context, req *videodemo.ChangeVideoRequest) error {
	resp, err := videoClient.IncCommentCount(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func DescVideoFavorite(ctx context.Context, req *videodemo.ChangeVideoRequest) error {
	resp, err := videoClient.DescFavoriteCount(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func DescVideoComment(ctx context.Context, req *videodemo.ChangeVideoRequest) error {
	resp, err := videoClient.DescCommentCount(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetVideoAuthor(ctx context.Context, req *videodemo.GetFavoriteVideoRequest) (int64, error) {
	resp, err := videoClient.GetFavoriteVideo(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	if len(resp.VideoList) == 0 {
		return 0, errno.ServiceErr
	}
	return resp.VideoList[0].Author.Id, nil
}
