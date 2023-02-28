package service

import (
	"context"
	"demo/tiktok/cmd/interact/dal/db"
	"demo/tiktok/cmd/interact/pack"
	"demo/tiktok/kitex_gen/interactdemo"
)

type GetFavoriteVideosService struct {
	ctx context.Context
}

func NewGetFavoriteVideosService(ctx context.Context) *GetFavoriteVideosService {
	return &GetFavoriteVideosService{ctx: ctx}
}

func (s *GetFavoriteVideosService) GetFavoriteVideos(req *interactdemo.GetUserFavoriteRequest) ([]*interactdemo.Video, error) {
	ids, err := db.GetFavoriteVideoIds(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return pack.VideoList(s.ctx, ids)
}
