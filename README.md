# IM聊天

即时聊天demo

- mysql 存储用户数据
- mongodb 存储聊天信息
- redis 缓存

## 第三方包
- [gin](https://github.com/gin-gonic/gin)
- [logrus](https://github.com/sirupsen/logrus)
- [gorm](https://github.com/go-gorm/gorm)
- [go-redis](https://github.com/go-redis/redis)
- [mongodb](https://github.com/mongodb/mongo-go-driver)
- [go-ini](https://github.com/go-ini/ini)
- [websocket](https://github.com/gorilla/websocket)
- stringer 通过go generate生成错误码文件

## 项目结构
```
.
├── README.md
├── api
├── cache
├── conf
├── config.ini
├── config.ini.example
├── go.mod
├── go.sum
├── main.go
├── model
├── pkg
├── router
├── serializer
└── service
```

## 功能
- 单对单聊天
- 在线、不在线应答
- 聊天记录查看