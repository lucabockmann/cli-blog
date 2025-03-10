package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// | ----------------------------
// | Funktion: GetEnvDataByName
// | ----------------------------
// | Read the .env File from root directory
// | You Pass the Variable Key and get the
// | Value of the ENV-Variable
// | ----------------------------
func GetEnvDataByName(key string) (string, error) {
	// | ----------------------------
	// | Load the .env File in the root directory
	// | gopath/backend/.env
	// | ----------------------------
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: ", err)
	}

	return os.Getenv(key), nil
}
