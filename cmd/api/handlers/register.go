package handlers

import (
	"context"
	"demo/tiktok/cmd/api/rpc"
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

var AuthMiddleware *jwt.HertzJWTMiddleware

func Register(ctx context.Context, c *app.RequestContext) {
	username := c.Query("username")
	password := c.Query("password")

	if len(username) == 0 || len(password) == 0 {
		register_sendResponse(c, errno.ParamErr, 0, "")
		return
	}

	uid, err := rpc.CreateUser(ctx, &userdemo.CreateUserRequest{
		UserName: username,
		Password: password,
	})
	if err != nil {
		register_sendResponse(c, err, uid, "")
		return
	}
	AuthMiddleware.LoginHandler(ctx, c)
}

func register_sendResponse(c *app.RequestContext, err error, user_id int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
		"user_id":     user_id,
		"token":       token,
	})
}
