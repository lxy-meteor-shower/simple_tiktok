package handlers

import (
	"context"
	"demo/tiktok/cmd/api/rpc"
	"demo/tiktok/kitex_gen/interactdemo"
	"demo/tiktok/kitex_gen/videodemo"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

func Comment(ctx context.Context, c *app.RequestContext) {
	videoId, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	comment_id, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
	comment_text := c.Query("comment_text")

	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))

	switch action_type {
	case 1:
		comment, err := rpc.CreateComment(ctx, &interactdemo.CreateCommentRequest{
			UserId:      userId,
			VideoId:     videoId,
			CommentText: comment_text,
		})
		if err != nil {
			comment_sendResponse(c, err, nil)
			return
		}

		err = rpc.IncVideoComment(ctx, &videodemo.ChangeVideoRequest{
			VideoId: videoId,
		})
		if err != nil {
			comment_sendResponse(c, err, nil)
			return
		}

		comment_sendResponse(c, errno.Success, comment)
	case 2:
		comment, err := rpc.DeleteComment(ctx, &interactdemo.DeleteCommentRequest{
			UserId:    userId,
			VideoId:   videoId,
			CommentId: comment_id,
		})
		if err != nil {
			comment_sendResponse(c, err, nil)
			return
		}

		err = rpc.DescVideoComment(ctx, &videodemo.ChangeVideoRequest{
			VideoId: videoId,
		})
		if err != nil {
			comment_sendResponse(c, err, nil)
			return
		}

		comment_sendResponse(c, errno.Success, comment)
	}
}

func comment_sendResponse(c *app.RequestContext, err error, comment *interactdemo.Comment) {

	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"comment":     comment,
	})
}
