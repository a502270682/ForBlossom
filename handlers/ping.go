package handlers

import (
	"context"
	"forBlossem/adapter/error_code"
	"forBlossem/officialaccount"
	"forBlossem/proto"
)

func PingHandler(ctx context.Context, req *proto.PingReq, rsp *proto.PingRsp) *error_code.ReplyError {
	rsp.Success = "hello"
	return nil
}

func AccessTokenGetHandler(ctx context.Context, req *struct{}, rsp *proto.AccessTokenRsp) *error_code.ReplyError {
	token, err := officialaccount.GetWechatAccount().GetAccessToken()
	if err != nil {
		return error_code.Error(error_code.CodeSystemError, err.Error())
	}
	rsp.Token = token
	return nil
}
