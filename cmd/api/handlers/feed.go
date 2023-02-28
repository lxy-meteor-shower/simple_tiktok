package handlers

import (
	"context"
	"demo/tiktok/cmd/api/rpc"
	"demo/tiktok/kitex_gen/videodemo"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

func Feed(ctx context.Context, c *app.RequestContext) {
	lt, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)

	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))

	timeObj := time.Unix(lt, 0)
	lastest_time := timeObj.Format("2006-01-02 15:04:05")

	videos, next_time, err := rpc.GetFeed(ctx, &videodemo.GetFeedRequest{
		LatestTime: lastest_time,
		UserId:     userId,
	})

	if err != nil {
		feed_sendResponse(c, err, time.Now().Unix(), nil)
		return
	}

	t, err := time.Parse("2006-01-02 15:04:05", next_time)
	if err != nil {
		feed_sendResponse(c, err, time.Now().Unix(), nil)
		return
	}
	feed_sendResponse(c, errno.Success, t.Unix(), videos)
}

func feed_sendResponse(c *app.RequestContext, err error, next_time int64, videos []*videodemo.Video) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"next_time":   next_time,
		"video_list":  videos,
	})
}
