package cache

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"log"
)

var (
	RedisClient *redis.Client

	// RedisConfig redis配置
	RedisConfig = new(redisConfig)
)

// redisConfig redis配置
type redisConfig struct {
	Host     string
	Port     int
	Password string
	Db       int
}

func init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("redis init load failed: %v", err)
	}

	loadRedis(file)
	Redis()
}

func loadRedis(file *ini.File) {
	section := file.Section("redis")
	RedisConfig.Host = section.Key("host").String()
	RedisConfig.Password = section.Key("password").String()

	port, err := section.Key("port").Int()
	if err != nil {
		log.Println(err)
	}
	RedisConfig.Port = port

	db, err := section.Key("db").Int()
	if err != nil {
		log.Println(err)
	}
	RedisConfig.Db = db
}

// Redis 连接redis
func Redis() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", RedisConfig.Host, RedisConfig.Port),
		Password: RedisConfig.Password,
		DB:       RedisConfig.Db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		logrus.Info(err)
		panic(err)
	}
	RedisClient = client
}
