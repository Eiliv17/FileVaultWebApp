package initializers

import (
	"log"
	"os"
)

func CreateFilesDir() {

	dirname := os.Getenv("FILE_DIR")

	// create directory for storing files
	err := os.MkdirAll("./"+dirname, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
