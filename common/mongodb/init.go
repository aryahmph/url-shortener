package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func New(user, password, host, port, name string) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	credential := options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		AuthSource:    name,
		Username:      user,
		Password:      password,
	}
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri).
		SetMinPoolSize(50).
		SetMaxPoolSize(100).
		SetMaxConnIdleTime(60 * time.Second).SetAuth(credential))
	if err != nil {
		panic(err)
	}

	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	return client.Database(name)
}
