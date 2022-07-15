package app

import (
	"context"
	"go-pos/helper"

	"go-pos/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func Connect() (*mongo.Database, error) {
	credential := options.Credential{
		AuthSource: config.AppConfig.MONGO_DB_AUTHENTICATION_DATABASE,
		Username:   config.AppConfig.MONGO_DB_USERNAME,
		Password:   config.AppConfig.MONGO_DB_PASSWORD,
	}

	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://" + config.AppConfig.MONGO_DB_HOST + ":" + config.AppConfig.MONGO_DB_PORT).SetAuth(credential)
	client, err := mongo.NewClient(clientOptions)

	helper.PanicIfError(err)

	err = client.Connect(ctx)

	helper.PanicIfError(err)

	return client.Database("yesdok"), nil
}
