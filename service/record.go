package service

import (
	"context"
	"gin-chat/conf"
)

// InsertMsg 聊天记录
// database 数据库
// collection 集合
// read 0未读 1已读
// expire 过期时间
func InsertMsg(database, collection, content string, read uint, expire int64) error {
	col := conf.MongoClient.Database(database).Collection(collection)

	_, err := col.InsertOne(context.TODO(), content)
	return err
}
