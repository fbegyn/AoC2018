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

	sol1 := prob1(numbers)
	fmt.Printf("Solution for prob1 is: %d\n", sol1)
	sol2 := prob2(numbers)
	fmt.Printf("Solution for prob2 is: %d\n", sol2)

}

func prob1(numbers []int) (freq int) {
	for _, v := range numbers {
		freq += v
	}
	return
}

func prob2(numbers []int) (freq int) {
	freqL := make(map[int]int)

	ok := true
	for ok {
		for _, v := range numbers {
			freq += v
			freqL[freq]++
			if freqL[freq] >= 2 {
				ok = false
				break
			}
		}
	}
	return
}
