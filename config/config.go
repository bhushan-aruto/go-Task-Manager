package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAdres  string
	DatabaseUrl  string
	DatabaseName string
}

func LoadConfig() *Config {

	serverMode := os.Getenv("SERVER_MODE")

	if serverMode == "dev" {
		if err := godotenv.Load(); err != nil {
			log.Fatalln("missing .env file. Error: ", err.Error())
		}
	} else if serverMode != "prod" {
		log.Println("Invalid SERVER_MODE. Please set SERVER_MODE to 'dev' or 'prod'.")

	}
	return &Config{

		DatabaseUrl:  getEnv("DATABASE_URL"),
		ServerAdres:  getEnv("SERVER_ADDRRESS"),
		DatabaseName: getEnv("DATABASE_NAME"),
	}

}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("missing or empty .env varible %v", key)
	}

	return value
}
