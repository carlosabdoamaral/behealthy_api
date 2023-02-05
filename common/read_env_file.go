package common

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariables() error {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		fmt.Println(errEnv.Error())
		fmt.Println("Failed to load env file. Make sure .env file exists!")
		return errEnv
	}

	DatabaseUser = os.Getenv("DB_USER")
	DatabasePass = os.Getenv("DB_PASS")
	DatabaseHost = os.Getenv("DB_HOST")
	DatabaseName = os.Getenv("DB_NAME")
	DatabasePort = os.Getenv("DB_PORT")
	DatabaseSSL = os.Getenv("DB_SSL")
	DatabaseDriver = os.Getenv("DB_DRIVER")

	RabbitURL = os.Getenv("RABBIT_URL")
	RabbitQueueName = os.Getenv("RABBIT_QUEUENAME")

	return nil
}
