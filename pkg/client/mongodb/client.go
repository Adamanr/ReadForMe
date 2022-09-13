package mongodb

import (
	logs "MyPay/pkg/logging"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func Mongo() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		logs.Logger().Fatalf("Не удалось подключиться к MongoDB по причине: %s", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Connect(ctx); err != nil {
		logs.Logger().Fatalf("Не удалось подключиться к MongoDB прошло время ожидания [%s]", err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			logs.Logger().Fatalf("Не удалось отключиться от MongoDB по причине: %s", err)
		}
	}()
	return client
}
