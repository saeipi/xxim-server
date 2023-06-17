// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package accountservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateRobotReq             = pb.CreateRobotReq
	CreateRobotResp            = pb.CreateRobotResp
	GetSelfUserInfoReq         = pb.GetSelfUserInfoReq
	GetSelfUserInfoResp        = pb.GetSelfUserInfoResp
	GetUserInfoReq             = pb.GetUserInfoReq
	GetUserInfoResp            = pb.GetUserInfoResp
	GetUserModelByIdReq        = pb.GetUserModelByIdReq
	GetUserModelByIdReq_Opt    = pb.GetUserModelByIdReq_Opt
	GetUserModelByIdResp       = pb.GetUserModelByIdResp
	RefreshUserAccessTokenReq  = pb.RefreshUserAccessTokenReq
	RefreshUserAccessTokenResp = pb.RefreshUserAccessTokenResp
	ResetUserAccountMapReq     = pb.ResetUserAccountMapReq
	ResetUserAccountMapResp    = pb.ResetUserAccountMapResp
	RevokeUserAccessTokenReq   = pb.RevokeUserAccessTokenReq
	RevokeUserAccessTokenResp  = pb.RevokeUserAccessTokenResp
	UpdateUserAccountMapReq    = pb.UpdateUserAccountMapReq
	UpdateUserAccountMapResp   = pb.UpdateUserAccountMapResp
	UpdateUserExtraMapReq      = pb.UpdateUserExtraMapReq
	UpdateUserExtraMapResp     = pb.UpdateUserExtraMapResp
	UpdateUserProfileMapReq    = pb.UpdateUserProfileMapReq
	UpdateUserProfileMapResp   = pb.UpdateUserProfileMapResp
	UserAccessTokenReq         = pb.UserAccessTokenReq
	UserAccessTokenResp        = pb.UserAccessTokenResp
	UserAfterOfflineReq        = pb.UserAfterOfflineReq
	UserAfterOfflineResp       = pb.UserAfterOfflineResp
	UserAfterOnlineReq         = pb.UserAfterOnlineReq
	UserAfterOnlineResp        = pb.UserAfterOnlineResp
	UserBeforeConnectReq       = pb.UserBeforeConnectReq
	UserBeforeConnectResp      = pb.UserBeforeConnectResp
	UserBeforeRequestReq       = pb.UserBeforeRequestReq
	UserBeforeRequestResp      = pb.UserBeforeRequestResp
	UserDestroyReq             = pb.UserDestroyReq
	UserDestroyResp            = pb.UserDestroyResp
	UserRegisterReq            = pb.UserRegisterReq
	UserRegisterResp           = pb.UserRegisterResp

	AccountService interface {
		// UserRegister 用户注册
		UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error)
		// UserDestroy 用户注销
		UserDestroy(ctx context.Context, in *UserDestroyReq, opts ...grpc.CallOption) (*UserDestroyResp, error)
		// UserAccessToken 用户登录
		UserAccessToken(ctx context.Context, in *UserAccessTokenReq, opts ...grpc.CallOption) (*UserAccessTokenResp, error)
		// RefreshUserAccessToken 刷新用户token
		RefreshUserAccessToken(ctx context.Context, in *RefreshUserAccessTokenReq, opts ...grpc.CallOption) (*RefreshUserAccessTokenResp, error)
		// RevokeUserAccessToken 注销用户token
		RevokeUserAccessToken(ctx context.Context, in *RevokeUserAccessTokenReq, opts ...grpc.CallOption) (*RevokeUserAccessTokenResp, error)
		// UpdateUserAccountMap 更新用户账号信息
		UpdateUserAccountMap(ctx context.Context, in *UpdateUserAccountMapReq, opts ...grpc.CallOption) (*UpdateUserAccountMapResp, error)
		// ResetUserAccountMap 重置用户账号信息
		ResetUserAccountMap(ctx context.Context, in *ResetUserAccountMapReq, opts ...grpc.CallOption) (*ResetUserAccountMapResp, error)
		// CreateRobot 创建机器人
		CreateRobot(ctx context.Context, in *CreateRobotReq, opts ...grpc.CallOption) (*CreateRobotResp, error)
	}

	defaultAccountService struct {
		cli zrpc.Client
	}
)

func NewAccountService(cli zrpc.Client) AccountService {
	return &defaultAccountService{
		cli: cli,
	}
}

// UserRegister 用户注册
func (m *defaultAccountService) UserRegister(ctx context.Context, in *UserRegisterReq, opts ...grpc.CallOption) (*UserRegisterResp, error) {
	client := pb.NewAccountServiceClient(m.cli.Conn())
	return client.UserRegister(ctx, in, opts...)
}

// UserDestroy 用户注销
func (m *defaultAccountService) UserDestroy(ctx context.Context, in *UserDestroyReq, opts ...grpc.CallOption) (*UserDestroyResp, error) {
	client := pb.NewAccountServiceClient(m.cli.Conn())
	return client.UserDestroy(ctx, in, opts...)
}

// UserAccessToken 用户登录
func (m *defaultAccountService) UserAccessToken(ctx context.Context, in *UserAccessTokenReq, opts ...grpc.CallOption) (*UserAccessTokenResp, error) {
	client := pb.NewAccountServiceClient(m.cli.Conn())
	return client.UserAccessToken(ctx, in, opts...)
}

// RefreshUserAccessToken 刷新用户token
func (m *defaultAccountService) RefreshUserAccessToken(ctx context.Context, in *RefreshUserAccessTokenReq, opts ...grpc.CallOption) (*RefreshUserAccessTokenResp, error) {
	client := pb.NewAccountServiceClient(m.cli.Conn())
	return client.RefreshUserAccessToken(ctx, in, opts...)
}

// RevokeUserAccessToken 注销用户token
func (m *defaultAccountService) RevokeUserAccessToken(ctx context.Context, in *RevokeUserAccessTokenReq, opts ...grpc.CallOption) (*RevokeUserAccessTokenResp, error) {
	client := pb.NewAccountServiceClient(m.cli.Conn())
	return client.RevokeUserAccessToken(ctx, in, opts...)
}

// UpdateUserAccountMap 更新用户账号信息
func (m *defaultAccountService) UpdateUserAccountMap(ctx context.Context, in *UpdateUserAccountMapReq, opts ...grpc.CallOption) (*UpdateUserAccountMapResp, error) {
	client := pb.NewAccountServiceClient(m.cli.Conn())
	return client.UpdateUserAccountMap(ctx, in, opts...)
}

// ResetUserAccountMap 重置用户账号信息
func (m *defaultAccountService) ResetUserAccountMap(ctx context.Context, in *ResetUserAccountMapReq, opts ...grpc.CallOption) (*ResetUserAccountMapResp, error) {
	client := pb.NewAccountServiceClient(m.cli.Conn())
	return client.ResetUserAccountMap(ctx, in, opts...)
}

// CreateRobot 创建机器人
func (m *defaultAccountService) CreateRobot(ctx context.Context, in *CreateRobotReq, opts ...grpc.CallOption) (*CreateRobotResp, error) {
	client := pb.NewAccountServiceClient(m.cli.Conn())
	return client.CreateRobot(ctx, in, opts...)
}
