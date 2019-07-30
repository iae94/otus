package config

import (
	"github.com/spf13/viper"
	"log"
)

type LoggerConfig struct {
	Level            string   `mapstructure:"level"`
	Encoding         string   `mapstructure:"encoding"`
	OutputPaths      []string `mapstructure:"outputPaths"`
	ErrorOutputPaths []string `mapstructure:"errorOutputPaths"`
}

type Config struct {
	Port uint
	LoggerConfig
}

func ReadConfig() (conf *Config, err error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("src/configs")
	viper.AddConfigPath(".")
	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Reading config error: %v \n", err)
		return nil, err
	}
	loggerConfig := LoggerConfig{}
	port := viper.GetUint("Port")
	err = viper.UnmarshalKey("Logger", &loggerConfig)
	if err != nil {
		log.Printf("Unmarshaling config error: %v \n", err)
		return nil, err
	}
	conf = &Config{port, loggerConfig}

	return conf, nil
}
