package util

import (
	"log"
	"os"
)

func ReadFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal("Could not read file")
	}
	return string(content)
}
