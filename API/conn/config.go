package conn

import (
	"log"
	"os"
    "github.com/joho/godotenv"
)

/*
*TESTED WORKING
loads .env file
differentiates between running in container or not by looking at value DOCKER_CONTAINER
returns .env variable
*/
func EnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file", err)
	} 
  
  	return os.Getenv(key)
}