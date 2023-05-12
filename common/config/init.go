package config

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

type (
	config struct {
		Port           int    `env:"PORT,unset" envDefault:"4001"`
		ZookeeperHosts string `env:"ZOOKEEPER_HOSTS,unset"`

		MongoDB
	}

	MongoDB struct {
		MongoUser     string `env:"DATABASE_USER,unset"`
		MongoPassword string `env:"DATABASE_PASSWORD,unset"`
		MongoHost     string `env:"DATABASE_HOST,unset"`
		MongoPort     string `env:"DATABASE_PORT,unset"`
		MongoName     string `env:"DATABASE_NAME,unset"`
	}
)

func LoadConfig() *config {
	cfg := &config{}
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		panic(err)
	}
	return cfg
}
