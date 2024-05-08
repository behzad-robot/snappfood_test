package env

import "github.com/spf13/viper"

type Env struct {
	Port               int
	PostgresqlDatabase string
	RabbitMQConnection string
}

func CreateEnv() (*Env, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	result := &Env{}
	err = viper.Unmarshal(&result)
	return result, err
}
