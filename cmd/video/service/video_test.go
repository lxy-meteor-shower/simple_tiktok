package service

import (
	"context"
	"demo/tiktok/cmd/video/dal"
	"demo/tiktok/cmd/video/rpc"
	"demo/tiktok/kitex_gen/videodemo"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var s *CreateVideoService
var s1 *UserVideoListService
var s2 *FeedService
var s3 *IncVideoFavoriteService

func TestVideo(t *testing.T) {
	req := &videodemo.CreateVideoRequest{
		Title:    "a title",
		PlayUrl:  "playurl",
		CoverUrl: "coverurl",
		UserId:   1,
	}
	err := s.CreateVideo(req)
	assert.Equal(t, nil, err)

	req3 := &videodemo.ChangeVideoRequest{
		VideoId: 1,
	}
	err = s3.IncVideoFavorite(req3)
	assert.Equal(t, nil, err)

	req1 := &videodemo.GetUserVideoRequest{
		UserId: 1,
	}
	videos, err := s1.UserVideoList(req1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(videos))
	assert.Equal(t, "a title", videos[0].Title)
	assert.Equal(t, int64(1), videos[0].FavoriteCount)

	req2 := &videodemo.GetFeedRequest{
		LatestTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	_, videos, err = s2.Feed(req2)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, len(videos))
	assert.Equal(t, "a title", videos[0].Title)
}

func TestMain(m *testing.M) {
	dal.Init()
	rpc.Init()
	s = NewCreateVideoService(context.Background())
	s1 = NewUserVideoListService(context.Background())
	s2 = NewFeedService(context.Background())
	s3 = NewIncVideoFavoriteService(context.Background())
	code := m.Run()
	os.Exit(code)
}
