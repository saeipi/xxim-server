package handler

import (
	"context"
	"errors"
	gatewayservicelogic "github.com/cherish-chat/xxim-server/app/gateway/internal/logic/gatewayservice"
	"github.com/cherish-chat/xxim-server/app/gateway/internal/svc"
	"github.com/cherish-chat/xxim-server/common/i18n"
	"github.com/cherish-chat/xxim-server/common/pb"
	"github.com/cherish-chat/xxim-server/common/utils"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"strings"
)

func UnifiedHandleHttp[REQ ReqInterface, RESP RespInterface](
	svcCtx *svc.ServiceContext,
	ctx *gin.Context,
	route Route[REQ, RESP],
) (REQ, *pb.GatewayApiResponse, error) {
	request := route.RequestPool.NewRequest()
	do := route.Do
	// 请求体中的数据 反序列化到 request 中
	contentType := ctx.ContentType()
	encoding := pb.EncodingProto_PROTOBUF
	// 判断是json还是protobuf
	apiRequest := &pb.GatewayApiRequest{}
	if strings.Contains(contentType, "application/json") {
		// json
		err := ctx.ShouldBindJSON(apiRequest)
		if err != nil {
			return request, &pb.GatewayApiResponse{
				Header: i18n.NewInvalidDataError(err.Error()),
				Body:   nil,
			}, err
		}
		err = utils.Json.Unmarshal(apiRequest.Body, request)
		if err != nil {
			return request, &pb.GatewayApiResponse{
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Header:    i18n.NewInvalidDataError(err.Error()),
				Body:      nil,
			}, err
		}
		request.SetHeader(apiRequest.Header)
		encoding = pb.EncodingProto_JSON
	} else if strings.Contains(contentType, "application/x-protobuf") {
		// protobuf
		body, _ := ctx.GetRawData()
		err := proto.Unmarshal(body, apiRequest)
		if err != nil {
			return request, &pb.GatewayApiResponse{
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Header:    i18n.NewInvalidDataError(err.Error()),
				Body:      nil,
			}, err
		}
		err = proto.Unmarshal(apiRequest.Body, request)
		if err != nil {
			return request, &pb.GatewayApiResponse{
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Header:    i18n.NewInvalidDataError(err.Error()),
				Body:      nil,
			}, err
		}
		request.SetHeader(apiRequest.Header)
	} else {
		return request, &pb.GatewayApiResponse{
			RequestId: apiRequest.RequestId,
			Path:      apiRequest.Path,
			Header:    i18n.NewInvalidDataError("invalid content type, please use application/json or application/x-protobuf"),
			Body:      nil,
		}, nil
	}
	requestHeader := request.GetHeader()
	if requestHeader == nil {
		return request, &pb.GatewayApiResponse{
			RequestId: apiRequest.RequestId,
			Path:      apiRequest.Path,
			Header:    i18n.NewInvalidDataError("invalid request header"),
			Body:      nil,
		}, nil
	}
	requestHeader.ClientIp = utils.Http.GetClientIP(ctx.Request)
	requestHeader.Encoding = encoding
	requestHeader.GatewayPodIp = utils.GetPodIp()

	// beforeRequest
	userBeforeRequestResp, err := svcCtx.CallbackService.UserBeforeRequest(ctx.Request.Context(), &pb.UserBeforeRequestReq{
		Header: requestHeader,
		Path:   apiRequest.Path,
	})
	if err != nil {
		return request, &pb.GatewayApiResponse{
			Header:    i18n.NewServerError(requestHeader, svcCtx.Config.Mode, err),
			RequestId: apiRequest.RequestId,
			Path:      apiRequest.Path,
			Body:      nil,
		}, err
	}

	var result *pb.GatewayApiResponse
	body, _ := proto.Marshal(userBeforeRequestResp)
	if len(body) > 0 {
		responseHeader := userBeforeRequestResp.GetHeader()
		if responseHeader != nil && responseHeader.Code != pb.ResponseCode_SUCCESS {
			return request, &pb.GatewayApiResponse{
				Header:    responseHeader,
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Body:      MarshalResponse(requestHeader, userBeforeRequestResp),
			}, nil
		}
	}

	requestHeader.UserId = userBeforeRequestResp.UserId

	response, err := do(ctx.Request.Context(), request)
	body, _ = proto.Marshal(response)
	if len(body) > 0 {
		responseHeader := response.GetHeader()
		if responseHeader == nil {
			responseHeader = i18n.NewOkHeader()
		}
		result = &pb.GatewayApiResponse{
			Header:    responseHeader,
			RequestId: apiRequest.RequestId,
			Path:      apiRequest.Path,
			Body:      MarshalResponse(requestHeader, response),
		}
	} else {
		if err != nil {
			statusErr, ok := status.FromError(err)
			if ok {
				err = errors.New(statusErr.Message())
			}
			result = &pb.GatewayApiResponse{
				Header:    i18n.NewServerError(requestHeader, svcCtx.Config.Mode, err),
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Body:      nil,
			}
		} else {
			result = &pb.GatewayApiResponse{
				Header:    i18n.NewOkHeader(),
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Body:      nil,
			}
		}
	}
	return request, result, err
}

func UnifiedHandleUniversal[REQ ReqInterface, RESP RespInterface](
	svcCtx *svc.ServiceContext,
	ctx context.Context,
	connection *gatewayservicelogic.UniversalConnection,
	apiRequest *pb.GatewayApiRequest,
	route Route[REQ, RESP],
) (REQ, *pb.GatewayApiResponse, error) {
	request := route.RequestPool.NewRequest()
	do := route.Do
	// 请求体中的数据 反序列化到 request 中
	// 判断是json还是protobuf
	requestHeader := apiRequest.Header
	if requestHeader.Encoding == pb.EncodingProto_JSON {
		err := utils.Json.Unmarshal(apiRequest.Body, request)
		if err != nil {
			return request, &pb.GatewayApiResponse{
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Header:    i18n.NewInvalidDataError(err.Error()),
				Body:      nil,
			}, err
		}
	} else if requestHeader.Encoding == pb.EncodingProto_PROTOBUF {
		err := proto.Unmarshal(apiRequest.Body, request)
		if err != nil {
			return request, &pb.GatewayApiResponse{
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Header:    i18n.NewInvalidDataError(err.Error()),
				Body:      nil,
			}, err
		}
	} else {
		return request, &pb.GatewayApiResponse{
			Header: i18n.NewInvalidDataError("invalid content type, please use application/json or application/x-protobuf"),
			Body:   nil,
		}, nil
	}
	request.SetHeader(requestHeader)
	if requestHeader == nil {
		return request, &pb.GatewayApiResponse{
			Header:    i18n.NewInvalidDataError("invalid request header"),
			Body:      nil,
			RequestId: apiRequest.RequestId,
			Path:      apiRequest.Path,
		}, nil
	}

	//beforeRequest
	userBeforeRequestResp, err := svcCtx.CallbackService.UserBeforeRequest(ctx, &pb.UserBeforeRequestReq{
		Header: requestHeader,
		Path:   apiRequest.Path,
	})
	if err != nil {
		return request, &pb.GatewayApiResponse{
			Header:    i18n.NewServerError(requestHeader, svcCtx.Config.Mode, err),
			RequestId: apiRequest.RequestId,
			Path:      apiRequest.Path,
			Body:      nil,
		}, err
	}

	var result *pb.GatewayApiResponse
	body, _ := proto.Marshal(userBeforeRequestResp)
	if len(body) > 0 {
		responseHeader := userBeforeRequestResp.GetHeader()
		if responseHeader != nil && responseHeader.Code != pb.ResponseCode_SUCCESS {
			return request, &pb.GatewayApiResponse{
				Header:    responseHeader,
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Body:      MarshalResponse(requestHeader, userBeforeRequestResp),
			}, nil
		}
	}
	response, err := do(ctx, request)
	body, _ = proto.Marshal(response)
	if len(body) > 0 {
		responseHeader := response.GetHeader()
		if responseHeader == nil {
			responseHeader = i18n.NewOkHeader()
		}
		result = &pb.GatewayApiResponse{
			RequestId: apiRequest.RequestId,
			Path:      apiRequest.Path,
			Header:    responseHeader,
			Body:      MarshalResponse(requestHeader, response),
		}
	} else {
		if err != nil {
			statusErr, ok := status.FromError(err)
			if ok {
				err = errors.New(statusErr.Message())
			}
			result = &pb.GatewayApiResponse{
				Header:    i18n.NewServerError(requestHeader, svcCtx.Config.Mode, err),
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Body:      nil,
			}
		} else {
			result = &pb.GatewayApiResponse{
				Header:    i18n.NewOkHeader(),
				RequestId: apiRequest.RequestId,
				Path:      apiRequest.Path,
				Body:      nil,
			}
		}
	}
	return request, result, err
}

func MarshalResponse(requestHeader *pb.RequestHeader, data proto.Message) []byte {
	if requestHeader == nil {
		return nil
	}
	switch requestHeader.Encoding {
	case pb.EncodingProto_PROTOBUF:
		protobuf, _ := proto.Marshal(data)
		return protobuf
	case pb.EncodingProto_JSON:
		json, _ := utils.Json.Marshal(data)
		return json
	default:
		return nil
	}
}

func MarshalWriteData(requestHeader *pb.RequestHeader, data *pb.GatewayApiResponse) []byte {
	if requestHeader == nil {
		return nil
	}
	writeData := &pb.GatewayWriteDataContent{
		DataType: pb.GatewayWriteDataType_Response,
		Response: data,
		Message:  nil,
		Notice:   nil,
	}
	switch requestHeader.Encoding {
	case pb.EncodingProto_PROTOBUF:
		protobuf, _ := proto.Marshal(writeData)
		return protobuf
	case pb.EncodingProto_JSON:
		json, _ := utils.Json.Marshal(writeData)
		return json
	default:
		return nil
	}
}
