package Utilities

import (
	"log"
	"os"
	"path/filepath"
)

func CheckEnvVariables(envVariables ...string) error {

	for _, envVar := range envVariables {
		if os.Getenv(envVar) == "" {
			log.Fatal(envVar + " environment variable missing")
		}
	}

	return nil
}
func GetBasePath() string {
	ex, err := os.Executable()
	if err != nil {
		return "."
	}
	return filepath.Dir(ex)
}
