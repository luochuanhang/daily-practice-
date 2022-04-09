package main

import (
	"context"
	"fmt"
	"lianxi/90_MongoDB/model"
	"lianxi/90_MongoDB/util"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// func main() {
// 	var (
// 		client     = util.GetMgoCli()
// 		err        error
// 		collection *mongo.Collection
// 		result     *mongo.InsertManyResult
// 		id         primitive.ObjectID
// 	)
// 	collection = client.Database("go_db").Collection("test")

// 	//批量插入
// 	result, err = collection.InsertMany(context.TODO(), []interface{}{
// 		model.LogRecord{
// 			JobName: "job10",
// 			Command: "echo 1",
// 			Err:     "",
// 			Content: "1",
// 			Tp: model.TimePorint{
// 				StartTime: time.Now().Unix(),
// 				EndTime:   time.Now().Unix() + 10,
// 			},
// 		},
// 		model.LogRecord{
// 			JobName: "job10",
// 			Command: "echo 2",
// 			Err:     "",
// 			Content: "2",
// 			Tp: model.TimePorint{
// 				StartTime: time.Now().Unix(),
// 				EndTime:   time.Now().Unix() + 10,
// 			},
// 		},
// 		model.LogRecord{
// 			JobName: "job10",
// 			Command: "echo 3",
// 			Err:     "",
// 			Content: "3",
// 			Tp: model.TimePorint{
// 				StartTime: time.Now().Unix(),
// 				EndTime:   time.Now().Unix() + 10,
// 			},
// 		},
// 		model.LogRecord{
// 			JobName: "job10",
// 			Command: "echo 4",
// 			Err:     "",
// 			Content: "4",
// 			Tp: model.TimePorint{
// 				StartTime: time.Now().Unix(),
// 				EndTime:   time.Now().Unix() + 10,
// 			},
// 		},
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if result == nil {
// 		log.Fatal("result nil")
// 	}
// 	for _, v := range result.InsertedIDs {
// 		id = v.(primitive.ObjectID)
// 		fmt.Println("自增ID", id.Hex())
// 	}
// }
func main() {
	var (
		client     = util.GetMgoCli()
		err        error
		collection *mongo.Collection
		cursor     *mongo.Cursor
	)
	//2.选择数据库 my_db里的某个表
	collection = client.Database("go_db").Collection("test")

	//如果直接使用 LogRecord{JobName: "job10"}是查不到数据的，因为其他字段有初始值0或者“”
	cond := model.FindByJobName{JobName: "job10"}

	//按照jobName字段进行过滤jobName="job10",翻页参数0-2
	if cursor, err = collection.Find(context.TODO(), cond, options.Find().SetSkip(0), options.Find().SetLimit(2)); err != nil {
		fmt.Println(err)
		return
	}
	//延迟关闭游标
	defer func() {
		if err = cursor.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	//遍历游标获取结果数据
	for cursor.Next(context.TODO()) {
		var lr model.LogRecord
		//反序列化Bson到对象
		if cursor.Decode(&lr) != nil {
			fmt.Print(err)
			return
		}
		//打印结果数据
		fmt.Println(lr)
	}

	//这里的结果遍历可以使用另外一种更方便的方式：
	var results []model.LogRecord
	if err = cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		fmt.Println(result)
	}
}
