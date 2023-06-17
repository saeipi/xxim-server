// Code generated by goctl. DO NOT EDIT.
// Source: conversation.proto

package server

import (
	"context"

	"github.com/cherish-chat/xxim-server/app/conversation/internal/logic/friendservice"
	"github.com/cherish-chat/xxim-server/app/conversation/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"
)

type FriendServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedFriendServiceServer
}

func NewFriendServiceServer(svcCtx *svc.ServiceContext) *FriendServiceServer {
	return &FriendServiceServer{
		svcCtx: svcCtx,
	}
}

// FriendApply 添加好友
func (s *FriendServiceServer) FriendApply(ctx context.Context, in *pb.FriendApplyReq) (*pb.FriendApplyResp, error) {
	l := friendservicelogic.NewFriendApplyLogic(ctx, s.svcCtx)
	return l.FriendApply(in)
}