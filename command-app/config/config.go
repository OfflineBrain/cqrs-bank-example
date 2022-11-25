package config

import "github.com/spf13/viper"

type Config struct {
	ServiceName         string `mapstructure:"SERVICE_NAME"`
	ServerPort          int32  `mapstructure:"SERVER_PORT"`
	MongoDbHost         string `mapstructure:"MONGODB_HOST"`
	MongoDbPort         int32  `mapstructure:"MONGODB_PORT"`
	KafkaHost           string `mapstructure:"KAFKA_HOST"`
	KafkaPort           int32  `mapstructure:"KAFKA_PORT"`
	KafkaAccountTxTopic string `mapstructure:"KAFKA_ACCOUNT_TRANSACTION_TOPIC"`
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
