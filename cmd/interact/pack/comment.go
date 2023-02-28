package pack

import (
	"context"
	"demo/tiktok/cmd/interact/dal/db"
	"demo/tiktok/cmd/interact/rpc"
	"demo/tiktok/kitex_gen/interactdemo"
	"demo/tiktok/kitex_gen/userdemo"
)

func Comment(ctx context.Context, c *db.Comment) (*interactdemo.Comment, error) {
	if c == nil {
		return nil, nil
	}

	u, err := rpc.GetUserInfo(ctx, &userdemo.MGetUserRequest{
		UserId: c.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &interactdemo.Comment{
		Id: c.Id,
		User: &interactdemo.User{
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
		Content:    c.Content,
		CreateDate: c.CreateAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func CommentList(ctx context.Context, cs []*db.Comment) ([]*interactdemo.Comment, error) {
	res := make([]*interactdemo.Comment, len(cs))
	for i, value := range cs {
		c, err := Comment(ctx, value)
		if err != nil {
			return nil, err
		}
		res[i] = c
	}
	return res, nil
}
