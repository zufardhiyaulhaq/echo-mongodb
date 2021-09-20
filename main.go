package main

import (
	"github.com/google/uuid"
	"github.com/tidwall/evio"
	"github.com/zufardhiyaulhaq/echo-mongodb/pkg/settings"
)

func main() {
	var events evio.Events

	settings, err := settings.NewSettings()
	if err != nil {
		panic(err.Error())
	}

	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		key := uuid.NewString()
		value := string(in)

		out = []byte(value)
		out = []byte(key)

		return
	}

	if err := evio.Serve(events, "tcp://0.0.0.0:"+settings.RedisEventPort); err != nil {
		panic(err.Error())
	}
}
