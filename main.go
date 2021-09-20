package main

import (
	"context"

	"github.com/tidwall/evio"
	mongodb_client "github.com/zufardhiyaulhaq/echo-mongodb/pkg/mongodb"
	"github.com/zufardhiyaulhaq/echo-mongodb/pkg/settings"
	"github.com/zufardhiyaulhaq/echo-mongodb/pkg/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	var events evio.Events

	settings, err := settings.NewSettings()
	if err != nil {
		panic(err.Error())
	}

	client := mongodb_client.New(context.Background(), settings)

	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		value := string(in)

		docID := primitive.NewObjectID()
		echo := types.Echo{
			ID:   docID,
			Echo: value,
		}
		err := client.InsertEcho(echo)
		if err != nil {
			out = []byte(err.Error())
			return
		}

		getEcho, err := client.GetEcho(docID)
		if err != nil {
			out = []byte(err.Error())
			return
		}

		out = []byte(getEcho.Echo)
		return
	}

	if err := evio.Serve(events, "tcp://0.0.0.0:"+settings.MongoDBEventPort); err != nil {
		panic(err.Error())
	}

	defer func() {
		if err = client.Close(); err != nil {
			panic(err)
		}
	}()
}
