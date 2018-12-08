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

// Log hello
type Log struct {
	id        int
	timestamp time.Time
	activity  int
}

func main() {
	src, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer src.Close()

	logs := []Log{}

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		var y, m, d, h, min int
		var desc string
		var l Log

		if !strings.Contains(line, "#") {
			_, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d] %s",
				&y, &m, &d, &h, &min,
				&desc)
			l.id = -1
			checkError(err)
		} else {
			_, err := fmt.Sscanf(line, "[%d-%d-%d %d:%d] %s #%d",
				&y, &m, &d, &h, &min,
				&desc, &l.id)
			checkError(err)
		}

		l.timestamp = time.Date(y, time.Month(m), d, h, min, 0, 0, time.UTC)

		if desc == "Guard" {
			l.activity = 0
		} else if desc == "wakes" {
			l.activity = 2
		} else {
			l.activity = 1
		}
		logs = append(logs, l)
	}

	sort.Slice(logs, func(i, j int) bool { return logs[i].timestamp.Before(logs[j].timestamp) })

	for i := 1; i < len(logs); i++ {
		if logs[i].id == -1 {
			logs[i].id = logs[i-1].id
		}
	}

	fmt.Printf("Most asleep guard with min multiplied: %d\n", part1(logs))
	fmt.Printf("Most frequent guard with min multiplied: %d\n", part2(logs))

}

func part1(logs []Log) int {
	// Calcluate the numer of minutes slept for each guard
	noOfMinutesSlept := make(map[int]int)
	for i := 0; i < len(logs); i++ {
		if _, ok := noOfMinutesSlept[logs[i].id]; ok {
			if logs[i].activity == 2 {
				// time awake - time asleep
				noOfMinutesSlept[logs[i].id] += int(logs[i].timestamp.Sub(logs[i-1].timestamp).Minutes())
			}
		} else {
			noOfMinutesSlept[logs[i].id] = 0
		}
	}

	// Find the guard Id with the most minuites slept
	var hSleepDuration, hDurationGuardID int
	for currentID, currentMinutesSlept := range noOfMinutesSlept {
		if currentMinutesSlept > hSleepDuration {
			hDurationGuardID = currentID
			hSleepDuration = currentMinutesSlept
		}
	}

	// Fill out sleepSchedule with the frequency that
	// that minute the guard is asleep in
	var sleepSchedule [59]int
	var lastAsleep int
	for i := 0; i < len(logs); i++ {
		// Filter for corrrect guard
		if logs[i].id == hDurationGuardID {
			// Check if activity is wake up or fell asleep
			if logs[i].activity == 1 {
				lastAsleep = logs[i].timestamp.Minute()
			} else if logs[i].activity == 2 {
				for j := lastAsleep; j < logs[i].timestamp.Minute(); j++ {
					sleepSchedule[j]++
				}
			}
		}
	}

	// Find most frequent minute
	var mostFrequentMin, hFrequency int
	for i := 0; i < len(sleepSchedule); i++ {
		if sleepSchedule[i] > hFrequency {
			hFrequency = sleepSchedule[i]
			mostFrequentMin = i
		}
	}

	return hDurationGuardID * mostFrequentMin
}

// part2 calculates the most common minute asleep * guard ID
func part2(logs []Log) int {
	// Init noOfMinutesSlept with keys as guard ids and empty sleep schedules
	noOfMinutesSlept := make(map[int][]int)
	for i := 0; i < len(logs); i++ {
		if _, ok := noOfMinutesSlept[logs[i].id]; !ok {
			noOfMinutesSlept[logs[i].id] = make([]int, 59)
		}
	}

	// Fill in all guards sleeping schedule
	for currentGuardID, currentSleepSchedule := range noOfMinutesSlept {
		var lastAsleep int
		for i := 0; i < len(logs); i++ {
			// Filter for correct guard
			if logs[i].id == currentGuardID {
				// Check if activity is wake up or fell asleep
				if logs[i].activity == 1 {
					lastAsleep = logs[i].timestamp.Minute()
				} else if logs[i].activity == 2 {
					for j := lastAsleep; j < logs[i].timestamp.Minute(); j++ {
						currentSleepSchedule[j]++
					}
				}
			}
		}
	}

	// Find most frequent minute for any guard
	var mostFrequentMin, hFrequency, hGuardID int
	for currentGuardID, currentSleepSchedule := range noOfMinutesSlept {
		for i := 0; i < len(currentSleepSchedule); i++ {
			if currentSleepSchedule[i] > hFrequency {
				hFrequency = currentSleepSchedule[i]
				mostFrequentMin = i
				hGuardID = currentGuardID
			}
		}
	}

	return hGuardID * mostFrequentMin
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
