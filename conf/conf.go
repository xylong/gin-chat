package conf

import (
	"go.mongodb.org/mongo-driver/mongo"
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

// mysqlConfig 数据库配置
type mysqlConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// mongoConfig mongo数据库配置
type mongoConfig struct {
	Host     string
	Port     int
	Password string
	Database string
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
