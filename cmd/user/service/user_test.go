package service

import (
	"context"
	"demo/tiktok/cmd/user/dal"
	"demo/tiktok/kitex_gen/userdemo"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var s *CreateUserService
var s1 *CheckUserService
var s2 *IncUserFavoriteService
var s3 *MGetUserService
var s4 *IncUserFavoritedService
var s5 *IncUserVideoService

func TestUser(t *testing.T) {
	req := &userdemo.CreateUserRequest{
		UserName: "lxy",
		Password: "123456",
	}
	err := s.CreateUser(req)
	assert.Equal(t, nil, err)

	req1 := &userdemo.CheckUserRequest{
		UserName: "lxy",
		Password: "123456",
	}
	_, err = s1.CheckUser(req1)
	assert.Equal(t, nil, err)

	req2 := &userdemo.ChangeUserRequest{
		UserId: 1,
	}
	err = s2.IncUserFavorite(req2)
	assert.Equal(t, nil, err)
	err = s5.IncUserVideo(req2)
	assert.Equal(t, nil, err)
	err = s4.IncUserFavorited(req2)
	assert.Equal(t, nil, err)

	req3 := &userdemo.MGetUserRequest{
		UserId: 1,
	}
	user, err := s3.MGetUser(req3)
	assert.Equal(t, nil, err)

	assert.Equal(t, "lxy", user.Name)
	assert.Equal(t, int64(1), user.FavoriteCount)
	assert.Equal(t, int64(1), user.TotalFavorited)
	assert.Equal(t, int64(1), user.WorkCount)
}

func TestMain(m *testing.M) {
	dal.Init()
	s = NewCreateUserService(context.Background())
	s1 = NewCheckUserService(context.Background())
	s2 = NewIncUserFavoriteService(context.Background())
	s3 = NewMGetUserService(context.Background())
	s4 = NewIncUserFavoritedService(context.Background())
	s5 = NewIncUserVideoService(context.Background())
	code := m.Run()
	os.Exit(code)
}
