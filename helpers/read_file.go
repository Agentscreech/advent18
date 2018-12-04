package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//ReadFile reads a file
func ReadFile() {
	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		fmt.Printf("%d", i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
