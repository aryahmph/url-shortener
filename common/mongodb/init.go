package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"strings"
	"time"
)

func New(user, password, hostsPorts, name string, minPool, maxPool int, replicaSet string) *mongo.Database {
	var uriSb strings.Builder
	uriSb.WriteString(fmt.Sprintf("mongodb://%s:%s@%s", user, password, hostsPorts))
	opts := options.Client()

	// Replica set enable
	if replicaSet != "-" {
		opts.
			SetReplicaSet(replicaSet).
			SetReadPreference(readpref.Secondary()).
			// These two options should make mongodb fault-tolerant
			SetReadConcern(readconcern.Majority()).
			SetWriteConcern(writeconcern.New(writeconcern.WMajority()))
	} else {
		uriSb.WriteString("/" + name)
	}

	opts.ApplyURI(uriSb.String()).
		SetMinPoolSize(uint64(minPool)).
		SetMaxPoolSize(uint64(maxPool))

	client, err := mongo.NewClient(opts)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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
