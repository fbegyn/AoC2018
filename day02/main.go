package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	src, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer src.Close()

	var ids []string

	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Text()
		ids = append(ids, line)
	}

	prob1(ids)
	prob2(ids)
}

func prob1(ids []string) {
	twos := 0
	threes := 0

	for _, id := range ids {
		freq := make(map[rune]int)
		for _, l := range id {
			freq[l]++
		}

		t := 0
		th := 0
		for _, amount := range freq {
			switch amount {
			case 2:
				t++
			case 3:
				th++
			}
		}

		if t > 0 {
			twos++
		}
		if th > 0 {
			threes++
		}
	}

	checksum := twos * threes
	fmt.Printf("Checksum of the IDs: %d\n", checksum)
}

func prob2(ids []string) {
	found := false
	var idA string
	diff := []rune{}
	indexes := []int{}
	for ind := 0; ind < len(ids)-1; ind++ {
		for _, id := range ids[ind+1:] {
			diff, indexes = difference(ids[ind], id)
			if len(diff) == 1 {
				idA = id
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	//fmt.Printf("These ids differ by 1:\n%s\n%s\n", idA, idB)

	remaining := []rune{}
	for i, r := range idA {
		if i != indexes[0] {
			remaining = append(remaining, r)
		}
	}
	fmt.Printf("The remaining letters are:\n%s\n", string(remaining))
}

func difference(idA, idB string) ([]rune, []int) {
	diff := []rune{}
	indexes := []int{}
	for ind, r := range idB {
		if idB[ind] != idA[ind] {
			diff = append(diff, r)
			indexes = append(indexes, ind)
		}
	}
	return diff, indexes
}
