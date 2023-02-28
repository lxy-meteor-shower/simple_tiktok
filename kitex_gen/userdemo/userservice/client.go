// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userservice

import (
	"context"
	userdemo "demo/tiktok/kitex_gen/userdemo"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateUser(ctx context.Context, Req *userdemo.CreateUserRequest, callOptions ...callopt.Option) (r *userdemo.CreateUserResponse, err error)
	CheckUser(ctx context.Context, Req *userdemo.CheckUserRequest, callOptions ...callopt.Option) (r *userdemo.CheckUserResponse, err error)
	MGetUser(ctx context.Context, Req *userdemo.MGetUserRequest, callOptions ...callopt.Option) (r *userdemo.MGetUserResponse, err error)
	IncFavorited(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error)
	IncVideo(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error)
	IncFavorite(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error)
	DescFavorited(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error)
	DescFavorite(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, Req *userdemo.CreateUserRequest, callOptions ...callopt.Option) (r *userdemo.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, Req)
}

func (p *kUserServiceClient) CheckUser(ctx context.Context, Req *userdemo.CheckUserRequest, callOptions ...callopt.Option) (r *userdemo.CheckUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckUser(ctx, Req)
}

func (p *kUserServiceClient) MGetUser(ctx context.Context, Req *userdemo.MGetUserRequest, callOptions ...callopt.Option) (r *userdemo.MGetUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MGetUser(ctx, Req)
}

func (p *kUserServiceClient) IncFavorited(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IncFavorited(ctx, Req)
}

func (p *kUserServiceClient) IncVideo(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IncVideo(ctx, Req)
}

func (p *kUserServiceClient) IncFavorite(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.IncFavorite(ctx, Req)
}

func (p *kUserServiceClient) DescFavorited(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DescFavorited(ctx, Req)
}

func (p *kUserServiceClient) DescFavorite(ctx context.Context, Req *userdemo.ChangeUserRequest, callOptions ...callopt.Option) (r *userdemo.ChangeUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DescFavorite(ctx, Req)
}
