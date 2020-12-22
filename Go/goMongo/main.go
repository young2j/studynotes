package main

import (
	"context"
	"fmt"
	"log"
	"mongo-notes/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
    // ctx
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, config.ClientOpts)
	if err != nil {
		log.Fatal(err)
	}
	// collection
	collection := client.Database("test").Collection("person")

	// session 读取策略
	sessOpts := options.Session().SetDefaultReadConcern(readconcern.Majority())
	session, err := client.StartSession(sessOpts)
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(context.TODO())
    
	// transaction 读取优先级
	transacOpts := options.Transaction().SetReadPreference(readpref.Primary())
	// 插入一条记录、查找一条记录在同一个事务中
	result, err := session.WithTransaction(ctx, func(sessionCtx mongo.SessionContext) (interface{}, error) {
		// insert one
		insertOneResult, err := collection.InsertOne(sessionCtx, bson.M{"name": "无名小角色", "level": 5})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("inserted id:", insertOneResult.InsertedID)

		// find one
		var result struct {
			Name  string `bson:"name,omitempty"`
			Level int    `bson:"level,omitempty"`
		}
		singleResult := collection.FindOne(sessionCtx, bson.M{"name": "无名小角色"})
		if err = singleResult.Decode(&result); err != nil {
			return nil, err
		}

		return result, err

	}, transacOpts)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("find one result: %+v \n", result)

}