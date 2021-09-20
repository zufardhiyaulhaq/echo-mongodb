package mongodb_client

import (
	"context"
	"fmt"
	"log"

	"github.com/zufardhiyaulhaq/echo-mongodb/pkg/settings"
	"github.com/zufardhiyaulhaq/echo-mongodb/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBClient struct {
	Context  context.Context
	Client   *mongo.Client
	Settings settings.Settings
}

func New(context context.Context, settings settings.Settings) MongoDBClient {
	client, err := mongo.Connect(context, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", settings.MongoDBUser, settings.MongoDBPassword, settings.MongoDBHost, settings.MongoDBPort, settings.MongoDBDatabase)))
	if err != nil {
		log.Fatal(err)
	}

	return MongoDBClient{
		Context:  context,
		Client:   client,
		Settings: settings,
	}
}

func (m MongoDBClient) Close() error {
	return m.Client.Disconnect(m.Context)
}

func (m MongoDBClient) InsertEcho(echo types.Echo) error {
	collection := m.Client.Database(m.Settings.MongoDBDatabase).Collection("echo")

	_, err := collection.InsertOne(m.Context, echo)
	if err != nil {
		return err
	}

	return nil
}

func (m MongoDBClient) GetEcho(id primitive.ObjectID) (types.Echo, error) {
	var echo types.Echo

	collection := m.Client.Database(m.Settings.MongoDBDatabase).Collection("echo")
	err := collection.FindOne(m.Context, bson.M{"_id": id}).Decode(&echo)
	if err != nil {
		return echo, err
	}

	return echo, nil
}
