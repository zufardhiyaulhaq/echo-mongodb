package settings

import (
	"github.com/kelseyhightower/envconfig"
)

type Settings struct {
	EchoPort        string `envconfig:"ECHO_PORT" default:"5000"`
	HTTPPort        string `envconfig:"HTTP_PORT" default:"80"`
	MongoDBHost     string `envconfig:"MONGODB_HOST"`
	MongoDBPort     string `envconfig:"MONGODB_PORT"`
	MongoDBDatabase string `envconfig:"MONGODB_DATABASE"`
	MongoDBUser     string `envconfig:"MONGODB_USER"`
	MongoDBPassword string `envconfig:"MONGODB_PASSWORD"`
}

func NewSettings() (Settings, error) {
	var settings Settings

	err := envconfig.Process("", &settings)
	if err != nil {
		return settings, err
	}

	return settings, nil
}
