package main

import (
	"context"
	"fmt"
	"time"

	// driver
	_ "github.com/go-sql-driver/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var localMongoOpts = options.Client().
	SetConnectTimeout(10 * time.Second).
	SetHosts([]string{"127.0.0.1:27017"}).
	SetMaxPoolSize(5).
	SetMinPoolSize(1).
	SetDirect(true)

// Doc ...
type Doc struct {
	ID primitive.ObjectID `bson:"_id"`
	UpdatedAt time.Time `bson:"updatedAt"`
}


func main() {
	ctx := context.TODO()
	client, _ := mongo.Connect(ctx, localMongoOpts)
	collection := client.Database("anquan").Collection("time_test")

	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	timeNow := time.Now().In(cstSh)
	fmt.Println("timeNow:", timeNow)

	collection.InsertOne(ctx, bson.M{"updatedAt": timeNow})

	var results []Doc
	cursor, _ := collection.Find(ctx, bson.M{})
	_ = cursor.All(ctx,&results)
	for _, result := range results {
  	fmt.Println(result)
	}
}

