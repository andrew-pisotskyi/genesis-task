package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress     string `mapstructure:"SERVER_ADDRESS"`
	BtcUahApiURl      string `mapstructure:"BTC_UAH_API_URL"`
	EmailsDataFile    string `mapstructure:"EMAILS_DATA_FILE"`
	SmtpHost          string `mapstructure:"SMTP_HOST"`
	SmtpPort          string `mapstructure:"SMTP_PORT"`
	EmailFrom         string `mapstructure:"EMAIL_FROM"`
	EmailFromPassword string `mapstructure:"EMAIL_FROM_PASSWORD"`
}

func NewConfig() *Config {
	config := Config{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &config
}
