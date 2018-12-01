package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var numbers []int

	src, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("failed to read file into scanner: %v", err)
	}
	defer src.Close()

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		freqCh, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("Failed to parse line into numbber: %v", err)
		}
		numbers = append(numbers, freqCh)
	}

	freqL := make(map[int]int)
	ok := true
	first := true
	freq := 0

	for ok {
		for _, v := range numbers {
			freq += v
			freqL[freq]++
			if freqL[freq] >= 2 {
				ok = false
				fmt.Printf("Solution for problem 2 is %d\n", freq)
				break
			}
		}
		if first {
			fmt.Printf("Solution for problem 1 is %d\n", freq)
			first = false
		}
	}
}
