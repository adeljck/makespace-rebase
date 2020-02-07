package conf

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"time"
)

type Database struct {
	Mongo *mongo.Client
}

var MONGODB *Database

//初始化
func MongoInit() {
	MONGODB = &Database{
		Mongo: SetConnect(),
	}
	if err := MONGODB.Mongo.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Print(err)
		os.Exit(0)
	}
}

// 连接设置
func SetConnect() *mongo.Client {
	uri := os.Getenv("MONGO_URI")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetMaxPoolSize(500)) // 连接池
	if err != nil {
		fmt.Println(err)
	}
	return client
}
