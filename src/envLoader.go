package src

import (
	"os"
	"strings"
)

func envReplacer(value string) string {
	if len(value) < 3 {
		return value
	}

	if value[0:2] == "${" {
		value = value[2 : len(value)-1]
		valueSplit := strings.Split(value, ":")

		value = getEnv(valueSplit[0], strings.Join(valueSplit[1:], ""))
		return value
	}
	return value

}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

