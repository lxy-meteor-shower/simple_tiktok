package main

import (
	"context"
	"demo/tiktok/cmd/video/pack"
	"demo/tiktok/cmd/video/service"
	videodemo "demo/tiktok/kitex_gen/videodemo"
	"demo/tiktok/pkg/errno"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// CreateVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) CreateVideo(ctx context.Context, req *videodemo.CreateVideoRequest) (resp *videodemo.CreateVideoResponse, err error) {
	resp = new(videodemo.CreateVideoResponse)

	err = service.NewCreateVideoService(ctx).CreateVideo(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetUserVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetUserVideo(ctx context.Context, req *videodemo.GetUserVideoRequest) (resp *videodemo.GetUserVideoResponse, err error) {
	resp = new(videodemo.GetUserVideoResponse)

	videos, err := service.NewUserVideoListService(ctx).UserVideoList(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}

// GetFeed implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFeed(ctx context.Context, req *videodemo.GetFeedRequest) (resp *videodemo.GetFeedResponse, err error) {
	resp = new(videodemo.GetFeedResponse)

	next_time, videos, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	resp.NextTime = next_time
	return resp, nil
}

// IncFavoriteCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) IncFavoriteCount(ctx context.Context, req *videodemo.ChangeVideoRequest) (resp *videodemo.ChangeVideoResponse, err error) {
	resp = new(videodemo.ChangeVideoResponse)

	err = service.NewIncVideoFavoriteService(ctx).IncVideoFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// IncCommentCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) IncCommentCount(ctx context.Context, req *videodemo.ChangeVideoRequest) (resp *videodemo.ChangeVideoResponse, err error) {
	resp = new(videodemo.ChangeVideoResponse)

	err = service.NewIncVideoCommentService(ctx).IncVideoComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DescFavoriteCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DescFavoriteCount(ctx context.Context, req *videodemo.ChangeVideoRequest) (resp *videodemo.ChangeVideoResponse, err error) {
	resp = new(videodemo.ChangeVideoResponse)

	err = service.NewDescVideoFavoriteService(ctx).DescVideoFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DescCommentCount implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) DescCommentCount(ctx context.Context, req *videodemo.ChangeVideoRequest) (resp *videodemo.ChangeVideoResponse, err error) {
	resp = new(videodemo.ChangeVideoResponse)

	err = service.NewDescVideoCommentService(ctx).DescVideoComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetFavoriteVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetFavoriteVideo(ctx context.Context, req *videodemo.GetFavoriteVideoRequest) (resp *videodemo.GetFavoriteVideoResponse, err error) {
	resp = new(videodemo.GetFavoriteVideoResponse)

	videos, err := service.NewGetIdVideosService(ctx).GetIdVideos(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}
