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
	file, err := os.Open("../inputs/day3.txt")
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
	// tracker := map[string]bool{}
	for scanner.Scan() {
		//need to parse out the input in to variables

		firstCut := strings.Split(scanner.Text(), " ")
		secondCut := strings.Split(firstCut[2], ":")[0]
		startingX, _ := strconv.Atoi(strings.Split(secondCut, ",")[0])
		startingY, _ := strconv.Atoi(strings.Split(secondCut, ",")[1])
		placeX, _ := strconv.Atoi(strings.Split(firstCut[3], "x")[0])
		placeY, _ := strconv.Atoi(strings.Split(firstCut[3], "x")[1])

		//nested for loop starting at a grid index startingX and startingY, move placeX times, incrementing the index and then move to the next line and repeat placeY times.
		for i := 0; i < placeY; i++ {
			for j := 0; j < placeX; j++ {
				//if startingX+j, startingY+i not in tracker
				//if grid[startingX+j][startingY+i] == 1
				grid[startingX+j][startingY+i]++
				//then place in tracker
				//else increment
			}
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//return or print len(tracker)

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
