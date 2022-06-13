package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const Collection = "humans"

// MongoDB groups structs needed to mongodb.
type MongoDB struct {
	Collection *mongo.Collection
}

// ConnectionOpts groups options to a mongo connection.
type ConnectionOpts struct {
	Ctx      context.Context
	Uri      string
	Database string
}

// NewMongoDB creates a new instance of MongoDB.
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
