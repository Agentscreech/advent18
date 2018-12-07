package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("../inputs/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	twosCounter := 0
	threesCounter := 0
	for scanner.Scan() {
		var twos bool
		var threes bool
		letters := map[rune]int{}
		for _, letter := range scanner.Text() {
			if _, found := letters[letter]; !found {
				letters[letter] = 1
			} else {
				letters[letter]++
			}
		}
		for i := range letters {
			if letters[i] == 2 && !twos {
				twos = true
				twosCounter++
			}
			if letters[i] == 3 && !threes {
				threes = true
				threesCounter++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(twosCounter * threesCounter)

}
