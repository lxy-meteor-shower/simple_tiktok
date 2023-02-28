package service

import (
	"context"
	"demo/tiktok/cmd/interact/dal"
	"demo/tiktok/cmd/interact/rpc"
	"demo/tiktok/kitex_gen/interactdemo"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s *GetFavoriteVideosService

func TestFavorite(t *testing.T) {
	req := &interactdemo.GetUserFavoriteRequest{
		UserId: 1,
	}
	ids, err := s.GetFavoriteVideos(req)
	assert.Equal(t, nil, err)
	fmt.Printf("%+v", ids)
}

func TestMain(m *testing.M) {
	dal.Init()
	rpc.Init()
	s = NewGetFavoriteVideosService(context.Background())
	code := m.Run()
	os.Exit(code)
}
