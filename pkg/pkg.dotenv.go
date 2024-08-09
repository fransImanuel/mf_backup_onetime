package pkg

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GodotEnv(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Errorln("Error loading .env file")
	}

	env := make(chan string, 1)
	//fmt.Println(os.Getenv("GO_ENV"))

	if os.Getenv("GO_ENV") != "production" {
		godotenv.Load(filepath.Join(".env"))
		env <- os.Getenv(key)
	} else {
		env <- os.Getenv(key)
	}

	return <-env
}
