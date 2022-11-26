package config

import "github.com/spf13/viper"

type Config struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	KafkaHost   string `mapstructure:"KAFKA_HOST"`
	KafkaPort   int32  `mapstructure:"KAFKA_PORT"`
	KafkaTopic  string `mapstructure:"KAFKA_ACCOUNT_TRANSACTION_TOPIC"`
	PgHost      string `mapstructure:"PG_HOST"`
	PgPort      int32  `mapstructure:"PG_PORT"`
	PgUser      string `mapstructure:"PG_USER"`
	PgPassword  string `mapstructure:"PG_PASSWORD"`
	PgDatabase  string `mapstructure:"PG_DATABASE"`
	ServerPort  int32  `mapstructure:"SERVER_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("default")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
