package pack

import (
	"demo/tiktok/kitex_gen/userdemo"
	"demo/tiktok/pkg/errno"
	"errors"
)

func BuildBaseResp(err error) *userdemo.BaseResp {
	if err == nil {
		return baseResp(errno.Success)
	}

	// parse ErrNo err
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return baseResp(e)
	}
	// see other err as service error
	s := errno.ServiceErr.WithMessage(err.Error())
	return baseResp(s)
}

func baseResp(err errno.ErrNo) *userdemo.BaseResp {
	return &userdemo.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
	}
}
