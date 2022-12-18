package application

import (
	"log"
	"mta-hosting-optimizer/config"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Cfg config.Config

func goDotEnvVariable(key string) string {
	log.Println("Getting value for key = ", key)
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
func LoadConfiguration() config.Config {
	log.Println("Loading the .env file to read configuration")
	dotenv := goDotEnvVariable("X")
	url := goDotEnvVariable("URL")
	method := goDotEnvVariable("METHOD")
	X, err := strconv.ParseInt(dotenv, 10, 64)
	if err != nil {
		log.Fatalf(err.Error())
	}
	Cfg.X = X
	Cfg.Url = url
	Cfg.Method = method
	return Cfg
}
