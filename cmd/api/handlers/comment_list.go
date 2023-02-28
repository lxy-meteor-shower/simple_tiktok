package handlers

import (
	"context"
	"demo/tiktok/cmd/api/rpc"
	"demo/tiktok/kitex_gen/interactdemo"
	"demo/tiktok/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func CommentList(ctx context.Context, c *app.RequestContext) {
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)

	comment_list, err := rpc.GetVideoComment(ctx, &interactdemo.GetVideoCommentRequest{
		VideoId: video_id,
	})
	if err != nil {
		commentList_sendResponse(c, err, nil)
		return
	}
	commentList_sendResponse(c, errno.Success, comment_list)
}

func commentList_sendResponse(c *app.RequestContext, err error, comment_list []*interactdemo.Comment) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code":  Err.ErrCode,
		"status_msg":   Err.ErrMsg,
		"comment_list": comment_list,
	})
}
