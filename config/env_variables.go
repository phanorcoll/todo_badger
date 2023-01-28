package config

import (
	"fmt"

	"github.com/spf13/viper"
)

//envVars holds the list of environment variables
//to be used across the app
type envVars struct {
	//The port number for the http server
	PORT   string `mapstructure:"SERVER_PORT"`
  //The secret key for JWT
	SECRET_KEY string `mapstructure:"SECRET_KEY"`
}

var (
	//EnvVariables holds the variables to use
	EnvVariables envVars
	//fileName for the environment variables
	fileName = ".env"
)

func init() {
	//set default values in case the variable doesnt exists
	viper.SetDefault("PORT", "3000")
	viper.SetConfigFile(fileName)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %w", err))
	}

	err := viper.Unmarshal(&EnvVariables)
	if err != nil {
		panic(fmt.Errorf("Unable to unmarchal data %w", err))
	}
}
