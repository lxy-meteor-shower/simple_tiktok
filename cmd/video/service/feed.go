package service

import (
	"context"
	"demo/tiktok/cmd/video/dal/db"
	"demo/tiktok/cmd/video/pack"
	"demo/tiktok/kitex_gen/videodemo"
	"time"
)

type FeedService struct {
	ctx context.Context
}

func NewFeedService(ctx context.Context) *FeedService {
	return &FeedService{ctx: ctx}
}

func (s *FeedService) Feed(req *videodemo.GetFeedRequest) (string, []*videodemo.Video, error) {
	videos, err := db.GetFeed(s.ctx, req.LatestTime)
	if err != nil {
		return time.Now().Format("2006-01-02 15:04:05"), nil, err
	}
	var next_time string
	if len(videos) == 0 {
		next_time = time.Now().Format("2006-01-02 15:04:05")
	} else {
		next_time = videos[len(videos)-1].CreateAt.Format("2006-01-02 15:04:05")
	}

	vs, err := pack.VideoList(s.ctx, videos, false, req.UserId)
	if err != nil {
		return time.Now().Format("2006-01-02 15:04:05"), nil, err
	}
	return next_time, vs, nil
}
