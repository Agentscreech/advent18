package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

type entry struct {
	timestamp time.Time
	eventType string
}

type eventList []entry

type sleepTime struct {
	start, end int
}

type guard struct {
	id          string
	sleepRange  []sleepTime
	sleepTotal  int
	topMin      int
	topMinCount int
}

func main() {

	//open file
	file, err := os.Open("../inputs/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//custom time layout in go -> Month:1, day:2, hour:15 or 3, minute:4, second:5, year:6 or 2006, timezone:7
	var events eventList
	for scanner.Scan() {

		//parse data, first set is date/time, second is either Guard start or sleep/wake event

		entrySplit := strings.Split(scanner.Text(), "] ")
		timestamp, _ := time.Parse("[2006-1-2 15:04", entrySplit[0])
		parsed := entry{
			timestamp: timestamp,
			eventType: entrySplit[1],
		}
		events = append(events, parsed)

	}

	//sort times
	sort.Sort(events)

	//find which guard slept the most
	guards := buildGuardList(events)
	//get total time each guard slept
	var sleepMax int
	var sleepChamp guard
	for i, g := range guards {
		var totalSleep int
		//sum sleep delta to counter
		for _, s := range g.sleepRange {
			totalSleep = totalSleep + s.end - s.start
		}
		//add to guard property
		g.sleepTotal = totalSleep
		//did they sleep more than the current champ?
		if totalSleep > sleepMax {
			sleepMax = totalSleep
			sleepChamp = g
		}
		//update guard list
		guards[i] = g
	}

	//tally up all the minutes guard was asleep
	minTracker := make(map[int]int)
	for _, sleepMin := range sleepChamp.sleepRange {
		for i := sleepMin.start; i < sleepMin.end; i++ {
			minTracker[i]++
		}
	}
	//find which minute that guard slept the most
	var topDuration int
	var topMin int
	for i, min := range minTracker {
		if min > topDuration {
			topDuration = min
			topMin = i
		}

	}

	//return guard id * most min slept.
	fmt.Println("SleepChamp id: ", sleepChamp.id, " TopMin slept: ", topMin)

	//part 2

	//Of all guards, which guard is most frequently asleep on the same minute?
	var consistentGuard guard
	for i, g := range guards {
		minTracker := make(map[int]int)
		for _, sleepMin := range g.sleepRange {
			for j := sleepMin.start; j < sleepMin.end; j++ {
				minTracker[j]++
			}
		}
		var topDuration int
		var topMin int
		for k, min := range minTracker {
			if min > topDuration {
				topDuration = min
				topMin = k
			}

		}
		g.topMin = topMin
		g.topMinCount = topDuration
		if g.topMinCount > consistentGuard.topMinCount {
			consistentGuard = g
		}
		guards[i] = g

	}
	fmt.Println("Most consistent guard,", consistentGuard.id, " They sleep at", consistentGuard.topMin, "most often")

}

func buildGuardList(events eventList) map[string]guard {

	guardList := make(map[string]guard)

	var activeGuard guard
	var sleepCounter sleepTime

	for _, event := range events {
		if strings.Contains(event.eventType, "Guard") {
			//found guard entry.  update guard or create new one
			//reset sleep counter
			sleepCounter = sleepTime{}
			//get guard id from log
			id := strings.Split(event.eventType, " ")[1]

			if gid, ok := guardList[id]; ok {
				//seen guard before
				activeGuard = gid
				continue
			} else {
				//make new entry
				activeGuard = guard{
					id:         id,
					sleepRange: []sleepTime{},
				}
				//add to list
				guardList[id] = activeGuard
			}
		} else {
			//sleep event
			if strings.Contains(event.eventType, "falls") {
				//add start
				sleepCounter.start = event.timestamp.Minute()
			}

			if strings.Contains(event.eventType, "wakes") {
				//add end
				sleepCounter.end = event.timestamp.Minute()
				//add range
				activeGuard.sleepRange = append(activeGuard.sleepRange, sleepCounter)
				//update guard list
				guardList[activeGuard.id] = activeGuard
			}

		}
	}
	return guardList
}

//fulfill sort interface, sort by timestamp
func (l eventList) Len() int {
	return len(l)
}

func (l eventList) Less(i, j int) bool {
	iDate := l[i].timestamp
	jDate := l[j].timestamp
	return iDate.Sub(jDate) < 0
}

func (l eventList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
