package dbhelper

import (
	"context"

	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Mysql Purpose
func ConnectDbMySql() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:@(localhost)/tap?charset=utf8&parseTime=True&loc=Local")
	db.LogMode(true)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// End Of MySql Purpose
var ctx = context.Background()

// MongoDb Purpose
func ConnectDbMongo() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("spectrum"), nil
}

// End Of Mongodb Purpose
