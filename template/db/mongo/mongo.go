package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const ConnectionTimeout = 10 * time.Second

func Connect(cfg Config) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.URI))
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	// connect to the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), ConnectionTimeout)
		defer done()

		if err := client.Connect(ctx); err != nil {
			return nil, fmt.Errorf("failed to connect to mongoDB: %w", err)
		}
	}
	// ping the mongodb
	{
		ctx, done := context.WithTimeout(context.Background(), ConnectionTimeout)
		defer done()

		if err := client.Ping(ctx, readpref.Primary()); err != nil {
			return nil, fmt.Errorf("failed to ping mongoDB: %w", err)
		}
	}

	return client.Database(cfg.Name), nil
}
