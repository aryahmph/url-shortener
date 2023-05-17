package config

import (
	"github.com/caarlos0/env/v8"
	_ "github.com/joho/godotenv/autoload"
)

type (
	config struct {
		Port           int    `env:"PORT,unset" envDefault:"4001"`
		ZookeeperHosts string `env:"ZOOKEEPER_HOSTS,unset"`
		RedisAddress   string `env:"REDIS_ADDRESS,unset"`

		MongoDB
	}

	MongoDB struct {
		MongoUser        string `env:"DATABASE_USER,unset"`
		MongoPassword    string `env:"DATABASE_PASSWORD,unset"`
		MongoHostsPorts  string `env:"DATABASE_HOSTS_PORTS,unset"`
		MongoName        string `env:"DATABASE_NAME,unset"`
		MongoMinPoolSize int    `env:"DATABASE_MIN_POOL_SIZE,unset"`
		MongoMaxPoolSize int    `env:"DATABASE_MAX_POOL_SIZE,unset"`
		MongoReplicaSet  string `env:"DATABASE_REPLICA_SET,unset"`
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
