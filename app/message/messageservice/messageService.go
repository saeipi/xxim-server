// Code generated by goctl. DO NOT EDIT!
// Source: message.proto

package messageservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	MessageInsertReq  = pb.MessageInsertReq
	MessageInsertResp = pb.MessageInsertResp

	MessageService interface {
		// MessageInsert 插入消息
		MessageInsert(ctx context.Context, in *MessageInsertReq, opts ...grpc.CallOption) (*MessageInsertResp, error)
	}

	defaultMessageService struct {
		cli zrpc.Client
	}
)

func NewMessageService(cli zrpc.Client) MessageService {
	return &defaultMessageService{
		cli: cli,
	}
}

// MessageInsert 插入消息
func (m *defaultMessageService) MessageInsert(ctx context.Context, in *MessageInsertReq, opts ...grpc.CallOption) (*MessageInsertResp, error) {
	client := pb.NewMessageServiceClient(m.cli.Conn())
	return client.MessageInsert(ctx, in, opts...)
}
