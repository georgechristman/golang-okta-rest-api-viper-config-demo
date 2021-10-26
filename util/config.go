package util

import "github.com/spf13/viper"

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variables
type Config struct {
	OktaClientId string `mapstructure:"OKTA_CLIENT_ID"`
	OktaIssuer   string `mapstructure:"OKTA_ISSUER"`
	Audience     string `mapstructure:"OKTA_AUDIENCE"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
