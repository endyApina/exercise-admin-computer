package config

import (
	"log"

	"github.com/spf13/viper"
)

type Secrets struct {
	ServiceName         string `mapstructure:"service_name"`
	ServicePort         string `mapstructure:"service_port"`
	DatabaseHost        string `mapstructure:"database_host"`
	DatabasePort        string `mapstructure:"database_port"`
	DatabaseName        string `mapstructure:"database_name"`
	DatabasePassword    string `mapstructure:"database_password"`
	PulsarURL           string `mapstructure:"pulsar_url"`
	PulsarCert          string `mapstructure:"pulsar_cert"`
	HTTPPort            string `mapstructure:"HTTP_PORT"`
	MessagingServiceURL string `mapstructure:"messaging_service_url"`
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
