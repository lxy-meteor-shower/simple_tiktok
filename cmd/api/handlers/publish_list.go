package handlers

import (
	"context"
	"demo/tiktok/cmd/api/rpc"
	"demo/tiktok/kitex_gen/videodemo"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

func PublishList(ctx context.Context, c *app.RequestContext) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	if userId == 0 {
		claims := jwt.ExtractClaims(ctx, c)
		userId = int64(claims[constants.IdentityKey].(float64))
	}

	videos, err := rpc.GetUserVideoList(ctx, &videodemo.GetUserVideoRequest{
		UserId: userId,
	})
	if err != nil {
		publishList_sendResponse(c, err, nil)
		return
	}

	publishList_sendResponse(c, errno.Success, videos)
}

func publishList_sendResponse(c *app.RequestContext, err error, videos []*videodemo.Video) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"video_list":  videos,
	})
}
