package main

import (
	"context"
	"fmt"
	"lianxi/90_MongoDB/model"
	"lianxi/90_MongoDB/util"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	var (
		client     = util.GetMgoCli()
		err        error
		collection *mongo.Collection
		result     *mongo.InsertManyResult
		id         primitive.ObjectID
	)
	collection = client.Database("go_db").Collection("test")

	//批量插入
	result, err = collection.InsertMany(context.TODO(), []interface{}{
		model.LogRecord{
			JobName: "job10",
			Command: "echo 1",
			Err:     "",
			Content: "1",
			Tp: model.TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		model.LogRecord{
			JobName: "job10",
			Command: "echo 2",
			Err:     "",
			Content: "2",
			Tp: model.TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		model.LogRecord{
			JobName: "job10",
			Command: "echo 3",
			Err:     "",
			Content: "3",
			Tp: model.TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
		model.LogRecord{
			JobName: "job10",
			Command: "echo 4",
			Err:     "",
			Content: "4",
			Tp: model.TimePorint{
				StartTime: time.Now().Unix(),
				EndTime:   time.Now().Unix() + 10,
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if result == nil {
		log.Fatal("result nil")
	}
	for _, v := range result.InsertedIDs {
		id = v.(primitive.ObjectID)
		fmt.Println("自增ID", id.Hex())
	}
}
