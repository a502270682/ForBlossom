package handlers

import (
	"context"
	wechat_account "forBlossem/adapter/account"
	"forBlossem/adapter/error_code"
	"forBlossem/proto"
)

func PingHandler(ctx context.Context, req *proto.PingReq, rsp *proto.PingRsp) *error_code.ReplyError {
	rsp.Success = "hello"
	return nil
}

func AccessTokenGetHandler(ctx context.Context, req *struct{}, rsp *proto.AccessTokenRsp) *error_code.ReplyError {
	token, err := wechat_account.GetWechatAccount().GetAccessToken()
	if err != nil {
		return error_code.Error(error_code.CodeSystemError, err.Error())
	}
	rsp.Token = token
	return nil
}
