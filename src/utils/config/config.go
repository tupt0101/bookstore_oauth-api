package config

import "github.com/spf13/viper"

type Config struct {
	DBUsername string `mapstructure:"mysql_users_username"`
	DBPassword string `mapstructure:"mysql_users_password"`
	DBHost     string `mapstructure:"mysql_users_host"`
	DBSchema   string `mapstructure:"mysql_users_schema"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
