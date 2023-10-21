package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func clean(value string) string {
	cleaned := strings.TrimSpace(value)

	prefixRemoved, prefixFound := strings.CutPrefix(cleaned, "\"")
	allQuotesRemoved, suffixFound := strings.CutSuffix(prefixRemoved, "\"")

	if prefixFound && suffixFound {
		cleaned = allQuotesRemoved
	}

	return cleaned
}

func LoadDotEnv() {
	file, err := os.Open(".env")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()
		key, value, found := strings.Cut(t, "=")

		if found && string(t[0]) != "#" {
			cleanedValue := clean(value)
			os.Setenv(key, cleanedValue)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
