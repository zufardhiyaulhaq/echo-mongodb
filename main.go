package main

import (
	"context"
	"sync"

	"github.com/rs/zerolog/log"
	"github.com/zufardhiyaulhaq/echo-mongodb/pkg/settings"

	mongodb_client "github.com/zufardhiyaulhaq/echo-mongodb/pkg/mongodb"
)

func main() {
	settings, err := settings.NewSettings()
	if err != nil {
		panic(err.Error())
	}

	log.Info().Msg("creating mongodb client")
	client := mongodb_client.New(context.Background(), settings)

	wg := new(sync.WaitGroup)
	wg.Add(2)

	log.Info().Msg("starting server")
	server := NewServer(settings, client)

	go func() {
		log.Info().Msg("starting HTTP server")
		server.ServeHTTP()
		wg.Done()
	}()

	wg.Wait()
}
