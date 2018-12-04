package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//Open file and throw it in a scanner to iterate over
	file, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var boxes []string

	//put all the lines (boxes) in a slice
	for scanner.Scan() {
		boxes = append(boxes, scanner.Text())
	}

	//for each one, go through them all and check if the are off by only one

	for _, box1 := range boxes {
		for _, box2 := range boxes {

			var letters []string
			var letterMismatch int

			if box1 == box2 {
				continue
			}
			for i := range box1 {
				if string(box1[i]) != string(box2[i]) {
					letterMismatch++
				} else {
					letters = append(letters, string(box1[i]))
				}
			}

			if letterMismatch == 1 {
				fmt.Println(strings.Join(letters, ""))
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//final output
	// fmt.Println(finalTwo)

}
