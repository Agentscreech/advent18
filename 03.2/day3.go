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
	tracker := map[int]bool{}
	for scanner.Scan() {
		//need to parse out the input in to variables

		firstCut := strings.Split(scanner.Text(), " ")
		secondCut := strings.Split(firstCut[2], ":")[0]
		startingX, _ := strconv.Atoi(strings.Split(secondCut, ",")[0])
		startingY, _ := strconv.Atoi(strings.Split(secondCut, ",")[1])
		placeX, _ := strconv.Atoi(strings.Split(firstCut[3], "x")[0])
		placeY, _ := strconv.Atoi(strings.Split(firstCut[3], "x")[1])
		id, _ := strconv.Atoi(strings.Split(firstCut[0], "#")[1])
		//nested for loop starting at a grid index startingX and startingY, move placeX times, setting the id claming that space at the index and then move to the next line and repeat placeY times.

		tracker[id] = true
		for i := 0; i < placeY; i++ {
			for j := 0; j < placeX; j++ {
				//if the index is 0, put the id there.  If it's not 0, set the id that's there and the current id to false in the tracker map
				if grid[startingX+j][startingY+i] != 0 {
					tracker[id] = false
					tracker[grid[startingX+j][startingY+i]] = false
				} else {
					grid[startingX+j][startingY+i] = id
				}
			}
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//loop through the tracker and print the true values.  Should only be one result
	for id, val := range tracker {
		if val {
			fmt.Println(id)
		}
	}

}
