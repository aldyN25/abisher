package configs

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type ConfDB struct {
	Host     string
	Port     string
	Dbname   string
	Username string
	Password string
}

type Configs struct {
	Dbconfig ConfDB

	Host string
}

var configs *Configs
var lock = &sync.Mutex{}

// Config func to get env value
func Config() *Configs {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	if configs == nil {
		lock.Lock()

		configs = &Configs{
			Dbconfig: ConfDB{
				Host:     os.Getenv("DB_HOST"),
				Port:     os.Getenv("DB_PORT"),
				Dbname:   os.Getenv("DB_NAME"),
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
			},

			Host: os.Getenv("HOST"),
		}
		lock.Unlock()
	}
	return configs

	// // load .env file
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	fmt.Print("Error loading .env file")
	// }
	// // Return the value of the variable
	// return os.Getenv(key)
}
