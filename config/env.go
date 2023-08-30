package config

import (
	"fmt"
	"os"
)

var (
	PORT = getEnv("PORT", "5000")
	DB   = getEnv("DB", "gotodo.sqlite")
)

func getEnv(name string, fallback string) string {
	if value, exists := os.LookupEnv(name); exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
