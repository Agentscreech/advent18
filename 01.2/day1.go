package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../inputs/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	tally := 0
	watcher := map[int]bool{}
	for {
		file.Seek(0, 0)
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			i, _ := strconv.Atoi(scanner.Text())
			tally += i
			if _, found := watcher[tally]; !found {
				watcher[tally] = true
			} else {
				fmt.Println(tally)
				os.Exit(0)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}
