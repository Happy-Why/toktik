// Code generated by Kitex v0.6.2. DO NOT EDIT.

package videoservice

import (
	"context"
	"github.com/Happy-Why/toktik-rpc/kitex_gen/video"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	VideoFeed(ctx context.Context, Req *video.VideoFeedRequest, callOptions ...callopt.Option) (r *video.VideoFeedResponse, err error)
	VideoPublish(ctx context.Context, Req *video.VideoPublishRequest, callOptions ...callopt.Option) (r *video.VideoPublishResponse, err error)
	PublishList(ctx context.Context, Req *video.PublishListRequest, callOptions ...callopt.Option) (r *video.PublishListResponse, err error)
	FavoriteList(ctx context.Context, Req *video.FavoriteListRequest, callOptions ...callopt.Option) (r *video.FavoriteListResponse, err error)
	FavoriteAction(ctx context.Context, Req *video.FavoriteActionRequest, callOptions ...callopt.Option) (r *video.FavoriteActionResponse, err error)
	CommentAction(ctx context.Context, Req *video.CommentActionRequest, callOptions ...callopt.Option) (r *video.CommentActionResponse, err error)
	CommentList(ctx context.Context, Req *video.CommentListRequest, callOptions ...callopt.Option) (r *video.CommentListResponse, err error)
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
	return &kVideoServiceClient{
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

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) VideoFeed(ctx context.Context, Req *video.VideoFeedRequest, callOptions ...callopt.Option) (r *video.VideoFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoFeed(ctx, Req)
}

func (p *kVideoServiceClient) VideoPublish(ctx context.Context, Req *video.VideoPublishRequest, callOptions ...callopt.Option) (r *video.VideoPublishResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VideoPublish(ctx, Req)
}

func (p *kVideoServiceClient) PublishList(ctx context.Context, Req *video.PublishListRequest, callOptions ...callopt.Option) (r *video.PublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, Req)
}

func (p *kVideoServiceClient) FavoriteList(ctx context.Context, Req *video.FavoriteListRequest, callOptions ...callopt.Option) (r *video.FavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, Req)
}

func (p *kVideoServiceClient) FavoriteAction(ctx context.Context, Req *video.FavoriteActionRequest, callOptions ...callopt.Option) (r *video.FavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, Req)
}

func (p *kVideoServiceClient) CommentAction(ctx context.Context, Req *video.CommentActionRequest, callOptions ...callopt.Option) (r *video.CommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, Req)
}

func (p *kVideoServiceClient) CommentList(ctx context.Context, Req *video.CommentListRequest, callOptions ...callopt.Option) (r *video.CommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, Req)
}
