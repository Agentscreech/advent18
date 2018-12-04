package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	//need a 2d array, 1000x1000
	var grid [][]int

	for i := 0; i < 1000; i++ { //couldn't remember shorthand for this
		row := make([]int, 1000)
		grid = append(grid, row)
	}

	var counter int

	for scanner.Scan() {
		//need to parse out the input in to variables

		firstCut := strings.Split(scanner.Text(), " ")
		secondCut := strings.Split(firstCut[2], ":")[0]
		// TODO: it'd be nice if I found a way to split on , then have the results be stored as int in an array
		startingX, _ := strconv.Atoi(strings.Split(secondCut, ",")[0])
		startingY, _ := strconv.Atoi(strings.Split(secondCut, ",")[1])
		placeX, _ := strconv.Atoi(strings.Split(firstCut[3], "x")[0])
		placeY, _ := strconv.Atoi(strings.Split(firstCut[3], "x")[1])

		//nested for loop starting at a startingX and startingY, increment the array index placeX times and then move to the next line and repeat placeY times.
		for i := 0; i < placeY; i++ {
			for j := 0; j < placeX; j++ {
				grid[startingX+j][startingY+i]++
			}
		}
		//if the index is > 1, increment a counter
		//TODO Find a way to do this during the loop

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//loop through the grid, counting how many are greater than 1
	for i, row := range grid {
		for j := range row {
			if grid[j][i] > 1 {
				counter++
			}
		}
	}

	fmt.Println(counter)
}
