package main

import (
	"context"
	"demo/tiktok/cmd/api/handlers"
	"demo/tiktok/cmd/api/minio1"
	"demo/tiktok/cmd/api/rpc"
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
)

func Init() {
	rpc.InitRPC()
	minio1.Init()
}

func main() {
	Init()
	r := server.New(
		server.WithHostPorts("0.0.0.0:8080"),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(1024*1024*1024),
	)

	authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) {
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg
			default:
				return e.Error()
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(consts.StatusOK, map[string]interface{}{
				"status_code": errno.Success.ErrCode,
				"status_msg":  errno.Success.ErrMsg,
				"user_id":     0,
				"token":       token,
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"status_code": errno.AuthorizationFailedErr.ErrCode,
				"status_msg":  errno.AuthorizationFailedErr.ErrMsg,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			username := c.Query("username")
			password := c.Query("password")

			if len(username) == 0 || len(password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			return rpc.CheckUser(context.Background(), &userdemo.CheckUserRequest{UserName: username, Password: password})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt, param: token, form: token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"code":    errno.ServiceErrCode,
				"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		},
	)))

	handlers.AuthMiddleware = authMiddleware
	v1 := r.Group("/douyin")
	user1 := v1.Group("/user")
	user1.POST("/login/", authMiddleware.LoginHandler)
	user1.POST("/register/", handlers.Register)
	user1.Use(authMiddleware.MiddlewareFunc())
	user1.GET("/", handlers.UserInfo)

	feed1 := v1.Group("/feed")
	feed1.Use(authMiddleware.MiddlewareFunc())
	feed1.GET("/", handlers.Feed)

	publish1 := v1.Group("/publish")
	publish1.Use(authMiddleware.MiddlewareFunc())
	publish1.POST("/action/", handlers.Publish)
	publish1.GET("/list/", handlers.PublishList)

	favorite1 := v1.Group("/favorite")
	favorite1.Use(authMiddleware.MiddlewareFunc())
	favorite1.GET("/list/", handlers.FavoriteList)
	favorite1.POST("/action/", handlers.Favorite)

	comment1 := v1.Group("/comment")
	comment1.Use(authMiddleware.MiddlewareFunc())
	comment1.GET("/list/", handlers.CommentList)
	comment1.POST("/action/", handlers.Comment)

	r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no route")
	})
	r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
		c.String(consts.StatusOK, "no method")
	})
	r.Spin()
}
