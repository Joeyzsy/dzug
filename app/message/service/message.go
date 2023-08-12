package service

import (
	"context"
	"dzug/app/message/infra/db"
	pb "dzug/protos/message"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type MsgSvr struct {
	pb.UnimplementedDouyinMessageServiceServer
}

func (MsgSvr) CreateMessage(ctx context.Context, request *pb.CreateMessageReq) (*pb.CreateMessageResp, error) {

	uuid := uuid.New().ID()

	message := &db.Message{
		FromUserId:  request.UserId,
		ToUserId:    request.ToUserId,
		Contents:    request.Content,
		MessageUUID: int64(uuid),
		CreateTime:  time.Now().Unix(),
	}

	if err := db.CreateMessage(ctx, message); err != nil {
		fmt.Printf("Create message fail " + err.Error())
		return nil, err
	}

	return &pb.CreateMessageResp{
		BaseResp: &pb.BaseResp{
			StatusCode: 200,
			StatusMsg:  "调用成功，你成功发送了一条消息",
		},
	}, nil

	return nil, nil
}

func (MsgSvr) GetMessageList(ctx context.Context, request *pb.GetMessageListReq) (*pb.GetMessageListResp, error) {

	//userId := tokenToUserId(request.Token)
	userId := int64(123)
	msgs, err := db.GetMessageList(ctx, userId, request.ToUserId, time.Now().Unix())
	if err != nil {
		fmt.Printf("Get messages fail: " + err.Error())
		return nil, err
	}
	infos := messagesToInfo(msgs)
	return &pb.GetMessageListResp{
		BaseResp: &pb.BaseResp{
			StatusCode: 200,
			StatusMsg:  "调用成功，你成功查询了消息记录",
		},
		MessageInfos: infos,
	}, nil
}

func messagesToInfo(messages []*db.Message) []*pb.MessageInfo {
	infos := make([]*pb.MessageInfo, 0)
	for _, msg := range messages {
		infos = append(infos, &pb.MessageInfo{
			MessageId:  msg.MessageUUID,
			FromUserId: msg.FromUserId,
			ToUserId:   msg.ToUserId,
			Content:    msg.Contents,
			CreateTime: msg.CreateTime,
		})
	}
	return infos
}
