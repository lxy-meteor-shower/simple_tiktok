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

func initInteractRPC() {
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

func CreateFavorite(ctx context.Context, req *interactdemo.FavoriteRequest) error {
	resp, err := interactClient.CreateFavorite(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func DeleteFavorite(ctx context.Context, req *interactdemo.FavoriteRequest) error {
	resp, err := interactClient.DeleteFavorite(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

func GetUserFavorite(ctx context.Context, req *interactdemo.GetUserFavoriteRequest) ([]*interactdemo.Video, error) {
	resp, err := interactClient.GetUserFavorite(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.VideoList, nil
}

func CreateComment(ctx context.Context, req *interactdemo.CreateCommentRequest) (*interactdemo.Comment, error) {
	resp, err := interactClient.CreateComment(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Comment, nil
}

func DeleteComment(ctx context.Context, req *interactdemo.DeleteCommentRequest) (*interactdemo.Comment, error) {
	resp, err := interactClient.DeleteComment(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Comment, nil
}

func GetVideoComment(ctx context.Context, req *interactdemo.GetVideoCommentRequest) ([]*interactdemo.Comment, error) {
	resp, err := interactClient.GetVideoComment(ctx, req)
	if err != nil {
		panic(err)
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.CommentList, nil
}
