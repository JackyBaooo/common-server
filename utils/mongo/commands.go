package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertOne(collection string, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	return DB.Collection(collection).InsertOne(ctx, document, opts...)
}

func FindOne(collection string, filter interface{}, options ...*options.FindOneOptions) *mongo.SingleResult {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	return DB.Collection(collection).FindOne(ctx, filter, options...)
}

func DeleteOne(collection string, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), Timeout)
	defer cancel()
	return DB.Collection(collection).DeleteOne(ctx, filter, opts...)
}
