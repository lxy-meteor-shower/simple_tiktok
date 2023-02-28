package service

import (
	"context"
	"demo/tiktok/cmd/interact/dal/db"
	"demo/tiktok/kitex_gen/interactdemo"
)

type CheckFavoriteService struct {
	ctx context.Context
}

func NewCheckFavoriteService(ctx context.Context) *CheckFavoriteService {
	return &CheckFavoriteService{ctx: ctx}
}

func (s *CheckFavoriteService) CheckFavorite(req *interactdemo.CheckFavoriteRequest) (bool, error) {
	favorites, err := db.GetOneFavorite(s.ctx, req.UserId, req.VideoId)
	if err != nil {
		return false, err
	}

	if len(favorites) == 0 {
		return false, nil
	}

	return true, nil
}
