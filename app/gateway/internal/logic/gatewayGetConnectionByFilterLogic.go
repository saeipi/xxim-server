package logic

import (
	"context"

	"github.com/cherish-chat/xxim-server/app/gateway/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GatewayGetConnectionByFilterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGatewayGetConnectionByFilterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GatewayGetConnectionByFilterLogic {
	return &GatewayGetConnectionByFilterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GatewayGetConnectionByFilterLogic) GatewayGetConnectionByFilter(in *pb.GatewayGetConnectionByFilterReq) (*pb.GatewayGetConnectionByFilterResp, error) {
	if len(in.GetFilter().GetUserIds()) > 0 {
		wsConnections := WsManager.wsConnectionMap.GetByUserIds(in.GetFilter().GetUserIds())
		var resp = &pb.GatewayGetConnectionByFilterResp{}
		for _, wsConnection := range wsConnections {
			resp.Connections = append(resp.Connections, wsConnection.ToPb())
		}
		return resp, nil
	}
	return &pb.GatewayGetConnectionByFilterResp{}, nil
}