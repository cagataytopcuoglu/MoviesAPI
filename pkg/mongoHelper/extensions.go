package mongoHelper

import (
	"MovieAPI/internal/config"
	"MovieAPI/pkg/log"
	"MovieAPI/pkg/pagination"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDb(settings *config.MongoSettings) (db *mongo.Database, err error) {
	uri := settings.Uri

	log.Logger.Infof("Mongo:Connection Uri:%s", uri)
	clientOptions := options.
		Client().
		ApplyURI(uri)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Logger.Fatalf("Mongo: couldn't connect to mongo: %v", err)
		return db, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		log.Logger.Fatalf("Mongo: mongo client couldn't connect with background context: %v", err)
		return db, err
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Logger.Fatalf("Mongo: Client Ping error", err)
	}

	db = client.Database(settings.DatabaseName)
	return db, err
}

func CastToId(id string) bson.M {

	objectIDS, _ := primitive.ObjectIDFromHex(id)

	return bson.M{
		"_id": objectIDS,
	}
}

func SetFindOptions(pageOptions *pagination.Pages) (findOptions *options.FindOptions) {

	findOptions = options.Find()
	if pageOptions != nil {
		if pageOptions.Offset() > 0 {
			findOptions.SetSkip(int64(pageOptions.Offset()))
		}
		if pageOptions.Limit() > 0 {
			findOptions.SetLimit(int64(pageOptions.Limit()))
		}
	}
	return findOptions
}

func BuildQuery(params map[string]string) (query bson.M) {

	query = bson.M{}
	for field, value := range params {

		if len(value) == 0 {
			continue
		}
		if field == "_id" {
			objectIDS, _ := primitive.ObjectIDFromHex(value)
			query["_id"] = objectIDS
		} else {
			query[field] = value
		}
	}
	return query
}
