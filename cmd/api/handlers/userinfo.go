package handlers

import (
	"context"
	"demo/tiktok/cmd/api/rpc"
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

func UserInfo(ctx context.Context, c *app.RequestContext) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)

	if userId == 0 {
		claims := jwt.ExtractClaims(ctx, c)
		userId = int64(claims[constants.IdentityKey].(float64))
	}

	user, err := rpc.GetUserInfo(ctx, &userdemo.MGetUserRequest{
		UserId: userId,
	})

	if err != nil {
		userinfo_sendResponse(c, err, nil)
		return
	}

	userinfo_sendResponse(c, errno.Success, user)
}

func userinfo_sendResponse(c *app.RequestContext, err error, user *userdemo.User) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"user":        user,
	})
}
