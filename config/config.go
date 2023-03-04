package config

import (
	"log"

	"github.com/spf13/viper"
)

type Secrets struct {
	DatabaseURL  string `mapstructure:"database_url"`
	DatabaseName string `mapstructure:"database_name"`
	PulsarURL    string `mapstructure:"pulsar_url"`
	PulsarCert   string `mapstructure:"pulsar_cert"`
	HTTPPort     string `mapstructure:"HTTP_PORT"`
}

// LoadSecrets loads up secrets from the .env file once
// If an evn file is present, secrets will be loaded. else ignored
func LoadSecrets(path string) (secrets Secrets, err error) {
	//Load secrets
	viper.SetConfigName("local")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//config file not found
			log.Println("config file ok: ", ok)
			//TODO: set default values
		} else {
			//config file was found
			log.Println("internal error")
			//but internal or another error produced.
		}
	}

	err = viper.Unmarshal(&secrets)
	return
}
