package storage

import (
	"context"

	parser "github.com/phuckhoa33/web-crawler/internal/parsers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveToMongo(data []parser.StockData) error {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	collection := client.Database("stocks").Collection("data")

	for _, stock := range data {
		_, err := collection.InsertOne(context.TODO(), stock)
		if err != nil {
			return err
		}
	}
	return nil
}
