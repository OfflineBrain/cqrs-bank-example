package config

import "github.com/spf13/viper"

type Config struct {
	ServiceName string `mapstructure:"SERVICE_NAME"`
	ServerPort  int32  `mapstructure:"SERVER_PORT"`
	PgHost      string `mapstructure:"PG_HOST"`
	PgPort      int32  `mapstructure:"PG_PORT"`
	PgUser      string `mapstructure:"PG_USER"`
	PgPassword  string `mapstructure:"PG_PASSWORD"`
	PgDatabase  string `mapstructure:"PG_DATABASE"`
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
