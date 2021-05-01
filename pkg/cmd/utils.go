package cmd

import (
	"log"
	"os"
	"strconv"
)

func GetenvString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func GetenvInt(key string, defaultVal int) int {
	if val, ok := os.LookupEnv(key); ok {
		i, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		return i
	}
	return defaultVal
}

func GetenvBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		b, err := strconv.ParseBool(val)
		if err != nil {
			log.Fatal(err)
		}
		return b
	}
	return defaultVal
}
