package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Collection = "humans"

type MongoDB struct {
	Collection *mongo.Collection
}

type ConnectionOpts struct {
	Ctx      context.Context
	Uri      string
	Database string
}

func NewMongoDB(opts *ConnectionOpts) (*MongoDB, error) {
	options := options.Client()
	options.ApplyURI(opts.Uri)
	options.SetConnectTimeout(10 * time.Second)

	client, err := mongo.Connect(opts.Ctx, options)
	if err != nil {
		return nil, err
	}

	err = client.Ping(opts.Ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoDB{
		Collection: client.Database(opts.Database).Collection(Collection),
	}, nil
}
