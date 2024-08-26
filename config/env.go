package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	LocalServerPort string `mapstructure:"LOCAL_SERVER_PORT"`
	MongoDBURL      string `mapstructure:"MONGODB_URL"`
	JWTSecret       string `mapstructure:"JWT_SECRET"`
	DBName          string `mapstructure:"DB_NAME"`
	ContextTimeout  int    `mapstructure:"CONTEXT_TIMEOUT"` // Changed to int for easier time.Duration usage
	SMTPUsername    string `mapstructure:"SMTP_USERNAME"`
	SMTPPassword    string `mapstructure:"SMTP_PASSWORD"`
	SMTPHost        string `mapstructure:"SMTP_HOST"`
	SMTPPort        string `mapstructure:"SMTP_PORT"`
}

func LoadEnv() *Env {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}

	var env Env

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatalf("Error unmarshalling .env file: %s", err)
	}

	return &env
}
