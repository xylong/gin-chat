package conf

import (
	"context"
	"fmt"
	"gin-chat/model"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/ini.v1"
	"log"
	"os"
)

var (
	MongoClient *mongo.Client

	// ServiceConfig 服务配置
	ServiceConfig = new(serviceConfig)

	// MysqlConfig 数据库配置
	MysqlConfig = new(mysqlConfig)

	// MongoConfig mongo数据库配置
	MongoConfig = new(mongoConfig)
)

// serviceConfig 服务配置
type serviceConfig struct {
	AppMode  string
	HttpPort int
}

func Init() {
	file, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	loadServer(file)
	loadMysql(file)
	loadMongo(file)

	Mongo()
	model.Database(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlConfig.User,
		MysqlConfig.Password,
		MysqlConfig.Host,
		MysqlConfig.Port,
		MysqlConfig.Database,
	))
}

func loadServer(file *ini.File) {
	section := file.Section("service")
	ServiceConfig.AppMode = section.Key("app_mode").String()

	port, err := section.Key("http_port").Int()
	if err != nil {
		log.Println(err)
	}
	ServiceConfig.HttpPort = port
}

// mysqlConfig 数据库配置
type mysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func loadMysql(file *ini.File) {
	section := file.Section("mysql")
	MysqlConfig.Host = section.Key("host").String()
	MysqlConfig.User = section.Key("user").String()
	MysqlConfig.Password = section.Key("password").String()
	MysqlConfig.Database = section.Key("database").String()

	port, err := section.Key("port").Int()
	if err != nil {
		log.Println(err)
	}
	MysqlConfig.Port = port
}

// mongoConfig mongo数据库配置
type mongoConfig struct {
	Host     string
	Port     int
	Password string
	Database string
}

func loadMongo(file *ini.File) {
	section := file.Section("mongo")
	MongoConfig.Host = section.Key("host").String()
	MongoConfig.Password = section.Key("password").String()
	MysqlConfig.Database = section.Key("database").String()

	port, err := section.Key("port").Int()
	if err != nil {
		log.Println(err)
	}
	MongoConfig.Port = port
}

func Mongo() {
	var err error
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", MongoConfig.Host, MongoConfig.Port))
	MongoClient, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		logrus.Info(err)
		panic(err)
	}

	logrus.Info("mongodb connected.")
}
