package handlers

import (
	"context"
	"demo/tiktok/cmd/api/minio1"
	"demo/tiktok/cmd/api/rpc"
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/kitex_gen/videodemo"
	"demo/tiktok/pkg/constants"
	"demo/tiktok/pkg/errno"
	"os/exec"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/jwt"
	"github.com/minio/minio-go/v7"
)

func Publish(ctx context.Context, c *app.RequestContext) {
	file, err := c.FormFile("data")
	if err != nil {
		publish_sendResponse(c, err)
		return
	}
	title := c.PostForm("title")
	claims := jwt.ExtractClaims(ctx, c)
	userId := int64(claims[constants.IdentityKey].(float64))

	fileReader, err := file.Open()
	if err != nil {
		publish_sendResponse(c, err)
		return
	}
	defer fileReader.Close()

	object_name := strconv.FormatInt(time.Now().Unix(), 10) + strconv.FormatInt(userId, 10)
	play_name := object_name + ".mp4"
	_, err = minio1.MinioClient.PutObject(ctx, "videos", play_name, fileReader, file.Size, minio.PutObjectOptions{})
	if err != nil {
		publish_sendResponse(c, err)
		return
	}

	playurl := "http://" + constants.PlayURL + ":9000/videos/" + play_name

	cmd := exec.Command("ffmpeg", "-i", playurl, "-ss", "0", "-frames:v", "1", "-f", "image2pipe", "-vcodec", "png", "-")
	outPipe, err := cmd.StdoutPipe()
	if err != nil {
		publish_sendResponse(c, err)
		return
	}
	if err := cmd.Start(); err != nil {
		publish_sendResponse(c, err)
		return
	}
	cover_name := object_name + ".png"
	_, err = minio1.MinioClient.PutObject(ctx, "images", cover_name, outPipe, int64(-1), minio.PutObjectOptions{ContentType: "image/png"})
	if err != nil {
		publish_sendResponse(c, err)
		return
	}
	if err := cmd.Wait(); err != nil {
		publish_sendResponse(c, err)
		return
	}
	coverurl := "http://" + constants.PlayURL + ":9000/images/" + cover_name

	err = rpc.CreateVideo(ctx, &videodemo.CreateVideoRequest{
		UserId:   userId,
		PlayUrl:  playurl,
		CoverUrl: coverurl,
		Title:    title,
	})
	if err != nil {
		publish_sendResponse(c, err)
		return
	}

	err = rpc.IncVideo(ctx, &userdemo.ChangeUserRequest{
		UserId: userId,
	})
	if err != nil {
		publish_sendResponse(c, err)
		return
	}
	publish_sendResponse(c, errno.Success)
}

func publish_sendResponse(c *app.RequestContext, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(consts.StatusOK, map[string]interface{}{
		"status_code": Err.ErrCode,
		"status_msg":  Err.ErrMsg,
	})
}
