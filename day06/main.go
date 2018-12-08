package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type coord []int

func main() {
	src, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer src.Close()

	coords := [][]int{}
	var maxX, maxY int
	minX := 100000000
	minY := 100000000
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := stringsAtoi(strings.Split(scanner.Text(), ", "))
		// keep track of the largest values (bounding box)
		if line[0] > maxX {
			maxX = line[0] + 1
		}
		if line[1] > maxY {
			maxY = line[1] + 1
		}
		if line[1] < minX {
			minX = line[1]
		}
		if line[1] < minY {
			minY = line[1]
		}
		// store the coords
		coords = append(coords, line)
	}

	// create distances
	offset := 5
	grid := make([][]string, maxY+offset)
	sum := make([][]int, maxY+offset)
	distances := make([][]map[string]int, maxY+offset)
	//closest := make([][]string, maxY)
	for y := range distances {
		grid[y] = make([]string, maxX+offset)
		sum[y] = make([]int, maxX+offset)
		distances[y] = make([]map[string]int, maxX+offset)
		//closest[y] = make([]string, maxX)
		for x := range distances[y] {
			grid[y][x] = ""
			distances[y][x] = make(map[string]int)
		}
	}

	freq := make(map[string]int)

	// fill the distancess
	for ind, co := range coords {
		grid[co[1]][co[0]] = strconv.FormatInt(int64(ind), 10)
		for i := range distances {
			for j := range distances[i] {
				curLoc := []int{j, i}
				distances[i][j][strconv.FormatInt(int64(ind), 10)] = manDist(co, curLoc)
				sum[i][j] = sumDist(distances[i][j])
			}
		}
	}

	for i := range distances {
		for j := range distances[i] {
			if i < maxY || i > minY || j < maxX || j > minX {
				closest := min(distances[i][j])
				grid[i][j] = closest
				if closest == "D" {
					continue
				}
			}
		}
	}

	safeCoords := make(map[string][][]int)
	for y := range grid {
		for x := range grid[y] {
			if x > maxX || y > maxY || x < minX || y < minY {
				delete(freq, grid[y][x])
				continue
			}
			index := grid[y][x]
			freq[index]++
			safeCoords[index] = append(safeCoords[index], []int{x, y})
		}
	}

	desiredRegion := [][]int{}
	for _, v := range safeCoords {
		for _, c := range v {
			if sum[c[1]][c[0]] < 10000 {
				desiredRegion = append(desiredRegion, c)
			}
		}
	}
	fmt.Printf("%v\n", len(desiredRegion))

	var maxField int
	for _, v := range freq {
		if v > maxField {
			maxField = v
		}
	}
	fmt.Printf("Largest non-infinite field is: %d\n", maxField)
}

func stringsAtoi(str []string) []int {
	var r []int
	for _, v := range str {
		x, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		r = append(r, x)
	}
	return r
}

func manDist(x, y []int) int {
	deltax := abs(x[0] - y[0])
	deltay := abs(x[1] - y[1])
	return deltax + deltay
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func printGrid(g [][]string) {
	for i := range g {
		for j := range g[i] {
			fmt.Printf("%v ", g[i][j])
		}
		fmt.Println()
	}
}

func sumDist(m map[string]int) (sum int) {
	for _, v := range m {
		sum += v
	}
	return sum
}

func min(m map[string]int) (ind string) {
	min := 320000000
	for k, v := range m {
		if v == min {
			ind = "D"
		}
		if v < min {
			min = v
			ind = k
		}
	}
	return
}

func max(m map[string]int) (ind string) {
	max := 0
	for k, v := range m {
		if v > max {
			max = v
			ind = k
		}
	}
	return
}
