package helper

import (
	"bufio"
	"log"
	"os"
)

//ReadFile reads a file and returns a scanner to iterate over
func ReadFile(location string) (*bufio.Scanner, error) {
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()
	return scanner, nil
}
