package Utilities

import (
	"log"
	"os"
)

func CheckEnvVariables(envVariables ...string) error {

	for _, envVar := range envVariables {
		if os.Getenv(envVar) == "" {
			log.Fatal(envVar + " environment variable missing")
		}
	}

	return nil
}
