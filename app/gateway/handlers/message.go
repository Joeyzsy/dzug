package handlers

import (
	"dzug/app/gateway/rpc"
	pb "dzug/protos/message"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MessageChatList(ctx *gin.Context) {
	var msgListReq pb.GetMessageListReq
	if err := ctx.Bind(&msgListReq); err != nil {
		ctx.JSON(http.StatusBadRequest, pb.GetMessageListResp{
			BaseResp: &pb.BaseResp{
				StatusCode: 400,
				StatusMsg:  "参数错误",
			},
			MessageInfos: nil,
		})
		return
	}
	userResp, err := rpc.MessageChatList(ctx, &msgListReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.GetMessageListResp{
			BaseResp: &pb.BaseResp{
				StatusCode: 500,
				StatusMsg:  "RPC服务调用错误",
			},
			MessageInfos: nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, userResp)
}

func MessagePostAction(ctx *gin.Context) {
	var msgPostReq pb.CreateMessageReq
	if err := ctx.Bind(&msgPostReq); err != nil {
		ctx.JSON(http.StatusBadRequest, pb.CreateMessageResp{
			BaseResp: &pb.BaseResp{
				StatusCode: 400,
				StatusMsg:  "参数错误",
			},
		})
		return
	}
	userResp, err := rpc.MessagePostAction(ctx, &msgPostReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pb.CreateMessageResp{
			BaseResp: &pb.BaseResp{
				StatusCode: 500,
				StatusMsg:  "RPC服务调用错误",
			},
		})
		return
	}
	ctx.JSON(http.StatusOK, userResp)
}
