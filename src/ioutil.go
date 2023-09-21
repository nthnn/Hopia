package main

import (
	"bufio"
	"os"

	"github.com/magiconair/properties"
)

func loadPropertiesFile(fileName string) *properties.Properties {
	return properties.MustLoadFile(fileName, properties.UTF8)
}

func readLinesFromFile(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
