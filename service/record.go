package service

import (
	"context"
	"gin-chat/conf"
	"gin-chat/model/ws"
	"time"
)

// InsertMsg 聊天记录
// database 数据库
// id 集合
// read 0未读 1已读
// expire 过期时间
func InsertMsg(database, id, content string, read uint, expire int64) error {
	collection := conf.MongoClient.Database(database).Collection(id) // 指定集合(没有就自动创建)

	now := time.Now().Unix()
	message := ws.Message{
		Content:   content,
		StartTime: now,
		EndTime:   now + expire,
		Read:      read,
	}

	_, err := collection.InsertOne(context.TODO(), message)
	return err
}
