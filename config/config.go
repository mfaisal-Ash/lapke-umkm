package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// load prod.env file
	err := godotenv.Load("prod.env")
	if err != nil {
		fmt.Print("Error loading prod.env file")
	}
	// Return the value of the variable
	return os.Getenv(key)
}

var StringAPIKEY = os.Getenv("APIKEY")
