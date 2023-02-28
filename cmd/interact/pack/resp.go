package pack

import (
	"demo/tiktok/kitex_gen/interactdemo"
	"demo/tiktok/pkg/errno"
	"errors"
)

func BuildBaseResp(err error) *interactdemo.BaseResp {
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

func baseResp(err errno.ErrNo) *interactdemo.BaseResp {
	return &interactdemo.BaseResp{
		StatusCode:    err.ErrCode,
		StatusMessage: err.ErrMsg,
	}
}
