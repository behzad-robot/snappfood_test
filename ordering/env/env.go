package env

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Env struct {
	Port               int
	PostgresqlDatabase string
	RabbitMQConnection string
}

func CreateEnv() (*Env, error) {
	_, e := os.Stat(".env")
	if e == nil {
		fmt.Println("reading env from .env file")
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("Viper Error", err)
			return nil, err
		}
		os.Setenv("PORT", viper.GetString("PORT"))
		os.Setenv("POSTGRESQL_DATABASE", viper.GetString("POSTGRESQL_DATABASE"))
		os.Setenv("RABBITMQ_CONNECTION", viper.GetString("RABBITMQ_CONNECTION"))
	}
	// Read from environment variables
	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 3000 // default port
	}

	db := os.Getenv("POSTGRESQL_DATABASE")
	if db == "" {
		return nil, fmt.Errorf("POSTGRESQL_DATABASE environment variable not set")
	}

	rabbitMQ := os.Getenv("RABBITMQ_CONNECTION")
	if rabbitMQ == "" {
		return nil, fmt.Errorf("RABBITMQ_CONNECTION environment variable not set")
	}

	return &Env{
		Port:               port,
		PostgresqlDatabase: db,
		RabbitMQConnection: rabbitMQ,
	}, nil
}
