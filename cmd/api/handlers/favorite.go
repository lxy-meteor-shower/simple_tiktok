package handlers

import (
	"context"
	"demo/tiktok/cmd/api/rpc"
	"demo/tiktok/kitex_gen/interactdemo"
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/kitex_gen/videodemo"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

func Favorite(ctx context.Context, c *app.RequestContext) {
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)

	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))

	authorId, err := rpc.GetVideoAuthor(ctx, &videodemo.GetFavoriteVideoRequest{
		VideoIds: []int64{video_id},
	})
	if err != nil {
		favorite_sendResponse(c, err)
		return
	}

	switch action_type {
	case 1:
		err = rpc.CreateFavorite(ctx, &interactdemo.FavoriteRequest{
			UserId:  userId,
			VideoId: video_id,
		})
		if err != nil {
			favorite_sendResponse(c, err)
			return
		}

		err = rpc.IncFavorite(ctx, &userdemo.ChangeUserRequest{
			UserId: userId,
		})
		if err != nil {
			favorite_sendResponse(c, err)
			return
		}

		err = rpc.IncFavorited(ctx, &userdemo.ChangeUserRequest{
			UserId: authorId,
		})
		if err != nil {
			favorite_sendResponse(c, err)
			return
		}

		err = rpc.IncVideoFavorite(ctx, &videodemo.ChangeVideoRequest{
			VideoId: video_id,
		})
		if err != nil {
			favorite_sendResponse(c, err)
			return
		}

	case 2:
		err = rpc.DeleteFavorite(ctx, &interactdemo.FavoriteRequest{
			UserId:  userId,
			VideoId: video_id,
		})
		if err != nil {
			favorite_sendResponse(c, err)
			return
		}

		err = rpc.DescFavorite(ctx, &userdemo.ChangeUserRequest{
			UserId: userId,
		})
		if err != nil {
			favorite_sendResponse(c, err)
			return
		}

		err = rpc.DescFavorited(ctx, &userdemo.ChangeUserRequest{
			UserId: userId,
		})
		if err != nil {
			favorite_sendResponse(c, err)
			return
		}

		err = rpc.DescFavorited(ctx, &userdemo.ChangeUserRequest{
			UserId: authorId,
		})
		if err != nil {
			favorite_sendResponse(c, err)
			return
		}

		err = rpc.DescVideoFavorite(ctx, &videodemo.ChangeVideoRequest{
			VideoId: video_id,
		})
		if err != nil {
			favorite_sendResponse(c, err)
			return
		}
	}
	favorite_sendResponse(c, errno.Success)
}

func favorite_sendResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
	})
}
