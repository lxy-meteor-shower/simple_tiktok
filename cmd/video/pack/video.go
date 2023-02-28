package pack

import (
	"context"
	"demo/tiktok/cmd/video/dal/db"
	"demo/tiktok/cmd/video/rpc"
	"demo/tiktok/kitex_gen/interactdemo"
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/kitex_gen/videodemo"
)

func Video(ctx context.Context, v *db.Video, isfavo bool, uid int64) (*videodemo.Video, error) {
	if v == nil {
		return nil, nil
	}

	u, err := rpc.GetUserInfo(ctx, &userdemo.MGetUserRequest{
		UserId: v.Author,
	})
	if err != nil {
		return nil, err
	}

	var isfavorite bool
	if isfavo == true {
		isfavorite = isfavo
	} else {
		isfavorite, err = rpc.IsFavorite(ctx, &interactdemo.CheckFavoriteRequest{
			UserId:  uid,
			VideoId: v.Id,
		})
	}

	return &videodemo.Video{
		Id: v.Id,
		Author: &videodemo.User{
			Id:              u.Id,
			Name:            u.Name,
			FollowCount:     u.FollowCount,
			FollowerCount:   u.FollowerCount,
			IsFollow:        u.IsFollow,
			Avatar:          u.Avatar,
			BackgroundImage: u.BackgroundImage,
			Signature:       u.Signature,
			TotalFavorited:  u.TotalFavorited,
			WorkCount:       u.WorkCount,
			FavoriteCount:   u.FavoriteCount,
		},
		PlayUrl:       v.PlayURL,
		CoverUrl:      v.CoverURL,
		FavoriteCount: v.FavoriteCount,
		CommentCount:  v.CommentCount,
		IsFavorite:    isfavorite,
		Title:         v.Title,
	}, nil
}

func VideoList(ctx context.Context, vs []*db.Video, isfavo bool, uid int64) ([]*videodemo.Video, error) {
	res := make([]*videodemo.Video, len(vs))
	for i, value := range vs {
		v, err := Video(ctx, value, isfavo, uid)
		if err != nil {
			return nil, err
		}
		res[i] = v
	}
	return res, nil
}
