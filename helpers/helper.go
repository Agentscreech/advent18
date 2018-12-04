package helper

import (
	"bufio"
	"log"
	"os"
)

//ReadFile reads a file and returns a scanner to iterate over
func ReadFile(location string) *bufio.Scanner {
	file, err := os.Open(location)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return scanner
}
