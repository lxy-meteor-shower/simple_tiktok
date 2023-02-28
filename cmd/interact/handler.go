package main

import (
	"context"
	"demo/tiktok/cmd/interact/pack"
	"demo/tiktok/cmd/interact/service"
	interactdemo "demo/tiktok/kitex_gen/interactdemo"
	"demo/tiktok/pkg/errno"
)

// InteractServiceImpl implements the last service interface defined in the IDL.
type InteractServiceImpl struct{}

// CreateFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CreateFavorite(ctx context.Context, req *interactdemo.FavoriteRequest) (resp *interactdemo.FavoriteResponse, err error) {
	resp = new(interactdemo.FavoriteResponse)

	err = service.NewCreateFavoriteService(ctx).CreateFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// DeleteFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) DeleteFavorite(ctx context.Context, req *interactdemo.FavoriteRequest) (resp *interactdemo.FavoriteResponse, err error) {
	resp = new(interactdemo.FavoriteResponse)

	err = service.NewDeleteFavoriteService(ctx).DeleteFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	return resp, nil
}

// GetUserFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetUserFavorite(ctx context.Context, req *interactdemo.GetUserFavoriteRequest) (resp *interactdemo.GetUserFavoriteResponse, err error) {
	resp = new(interactdemo.GetUserFavoriteResponse)

	videos, err := service.NewGetFavoriteVideosService(ctx).GetFavoriteVideos(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.VideoList = videos
	return resp, nil
}

// CheckFavorite implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CheckFavorite(ctx context.Context, req *interactdemo.CheckFavoriteRequest) (resp *interactdemo.CheckFavoriteResponse, err error) {
	resp = new(interactdemo.CheckFavoriteResponse)

	is_favo, err := service.NewCheckFavoriteService(ctx).CheckFavorite(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.IsFavorite = is_favo
	return resp, nil
}

// CreateComment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) CreateComment(ctx context.Context, req *interactdemo.CreateCommentRequest) (resp *interactdemo.CommentResponse, err error) {
	resp = new(interactdemo.CommentResponse)

	comment, err := service.NewCreateCommentService(ctx).CreateComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Comment = comment
	return resp, nil
}

// DeleteComment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) DeleteComment(ctx context.Context, req *interactdemo.DeleteCommentRequest) (resp *interactdemo.CommentResponse, err error) {
	resp = new(interactdemo.CommentResponse)

	err = service.NewDeleteCommentService(ctx).DeleteComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.Comment = nil
	return resp, nil
}

// GetVideoComment implements the InteractServiceImpl interface.
func (s *InteractServiceImpl) GetVideoComment(ctx context.Context, req *interactdemo.GetVideoCommentRequest) (resp *interactdemo.GetVideoCommentRespense, err error) {
	resp = new(interactdemo.GetVideoCommentRespense)

	comments, err := service.NewGetVideoCommentService(ctx).GetVideoComment(req)
	if err != nil {
		resp.BaseResp = pack.BuildBaseResp(err)
		return resp, nil
	}
	resp.BaseResp = pack.BuildBaseResp(errno.Success)
	resp.CommentList = comments
	return resp, nil
}
