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

type Interface interface {
	Close() error
	InsertEcho(echo types.Echo) error
	GetEcho(id primitive.ObjectID) (types.Echo, error)
}

type Client struct {
	Context  context.Context
	Client   *mongo.Client
	Settings settings.Settings
}

func New(context context.Context, settings settings.Settings) Client {
	client, err := mongo.Connect(context, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", settings.MongoDBUser, settings.MongoDBPassword, settings.MongoDBHost, settings.MongoDBPort, settings.MongoDBDatabase)))
	if err != nil {
		log.Fatal(err)
	}

	return Client{
		Context:  context,
		Client:   client,
		Settings: settings,
	}
}

func (m Client) Close() error {
	return m.Client.Disconnect(m.Context)
}

func (m Client) InsertEcho(echo types.Echo) error {
	collection := m.Client.Database(m.Settings.MongoDBDatabase).Collection("echo")

	_, err := collection.InsertOne(m.Context, echo)
	if err != nil {
		return err
	}

	return nil
}

func (m Client) GetEcho(id primitive.ObjectID) (types.Echo, error) {
	var echo types.Echo

	collection := m.Client.Database(m.Settings.MongoDBDatabase).Collection("echo")
	err := collection.FindOne(m.Context, bson.M{"_id": id}).Decode(&echo)
	if err != nil {
		return echo, err
	}

	return echo, nil
}
