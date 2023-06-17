// Code generated by goctl. DO NOT EDIT.
// Source: third.proto

package captchaservice

import (
	"context"

	"github.com/cherish-chat/xxim-server/common/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CaptchaVerifyReq    = pb.CaptchaVerifyReq
	CaptchaVerifyResp   = pb.CaptchaVerifyResp
	EmailCodeSendReq    = pb.EmailCodeSendReq
	EmailCodeSendResp   = pb.EmailCodeSendResp
	EmailCodeVerifyReq  = pb.EmailCodeVerifyReq
	EmailCodeVerifyResp = pb.EmailCodeVerifyResp
	GetCaptchaReq       = pb.GetCaptchaReq
	GetCaptchaResp      = pb.GetCaptchaResp
	SmsCodeSendReq      = pb.SmsCodeSendReq
	SmsCodeSendResp     = pb.SmsCodeSendResp
	SmsCodeVerifyReq    = pb.SmsCodeVerifyReq
	SmsCodeVerifyResp   = pb.SmsCodeVerifyResp

	CaptchaService interface {
		// GetCaptcha 获取图形验证码
		GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error)
		// CaptchaVerify 验证图形验证码
		CaptchaVerify(ctx context.Context, in *CaptchaVerifyReq, opts ...grpc.CallOption) (*CaptchaVerifyResp, error)
	}

	defaultCaptchaService struct {
		cli zrpc.Client
	}
)

func NewCaptchaService(cli zrpc.Client) CaptchaService {
	return &defaultCaptchaService{
		cli: cli,
	}
}

// GetCaptcha 获取图形验证码
func (m *defaultCaptchaService) GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error) {
	client := pb.NewCaptchaServiceClient(m.cli.Conn())
	return client.GetCaptcha(ctx, in, opts...)
}

// CaptchaVerify 验证图形验证码
func (m *defaultCaptchaService) CaptchaVerify(ctx context.Context, in *CaptchaVerifyReq, opts ...grpc.CallOption) (*CaptchaVerifyResp, error) {
	client := pb.NewCaptchaServiceClient(m.cli.Conn())
	return client.CaptchaVerify(ctx, in, opts...)
}
