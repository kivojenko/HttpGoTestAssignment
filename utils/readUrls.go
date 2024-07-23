package utils

import (
	"HttpGoTestAssignment/configuration"
	"bufio"
	"fmt"
	"os"
)

func ReadURLsFromFile() []string {
	cfg := configuration.GetConfiguration()
	filename := cfg.UrlsFile

	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("failed to open file %s: %w", filename, err))
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Errorf("error reading file %s: %w", filename, err))
	}

	return urls
}
