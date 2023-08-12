package db

import (
	"context"
)

func CreateMessage(ctx context.Context, message *Message) error {

	err := DB.WithContext(ctx).Create(&message).Error

	return err
}

func GetMessageList(ctx context.Context, userId int64, toUserID int64, latestTime int64) ([]*Message, error) {
	var messages []*Message
	err := DB.WithContext(ctx).Where("from_user_id = ? AND to_user_id = ? AND create_time >= ?",
		userId, toUserID, latestTime).Find(&messages).Error
	if err != nil {
		return nil, err
	}
	return messages, nil
}
