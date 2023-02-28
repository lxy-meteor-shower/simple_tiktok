package pack

import (
	"context"
	"demo/tiktok/cmd/interact/rpc"
	"demo/tiktok/kitex_gen/interactdemo"
	"demo/tiktok/kitex_gen/videodemo"
)

func VideoList(ctx context.Context, ids []int64) ([]*interactdemo.Video, error) {
	if ids == nil || len(ids) == 0 {
		return nil, nil
	}

	videos, err := rpc.GetIdVideos(ctx, &videodemo.GetFavoriteVideoRequest{
		VideoIds: ids,
	})
	if err != nil {
		return nil, err
	}

	res := make([]*interactdemo.Video, len(ids))
	for i, v := range videos {
		res[i] = &interactdemo.Video{
			Id: v.Id,
			Author: &interactdemo.User{
				Id:              v.Author.Id,
				Name:            v.Author.Name,
				FollowCount:     v.Author.FollowCount,
				FollowerCount:   v.Author.FollowerCount,
				IsFollow:        v.Author.IsFollow,
				Avatar:          v.Author.Avatar,
				BackgroundImage: v.Author.BackgroundImage,
				Signature:       v.Author.Signature,
				TotalFavorited:  v.Author.TotalFavorited,
				WorkCount:       v.Author.WorkCount,
				FavoriteCount:   v.Author.FavoriteCount,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    v.IsFavorite,
			Title:         v.Title,
		}
	}
	return res, nil
}
