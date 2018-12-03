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
	src, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Failed to open file: %v\n", err)
	}
	defer src.Close()

	claims := make(map[int][][]int)
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		nr, err := strconv.Atoi(strings.TrimLeft(line[0], "#"))
		if err != nil {
			log.Fatal(err)
		}
		offset := conv(strings.Split(strings.TrimRight(line[2], ":"), ","))
		size := conv(strings.Split(line[3], "x"))
		claims[nr] = [][]int{offset, size}
	}

	matrix := [1024][1024]rune{}
	overlap := 0
	for _, cl := range claims {
		offx, offy := coord(cl[0])
		x, y := coord(cl[1])
		for i := offx; i < offx+x; i++ {
			for j := offy; j < offy+y; j++ {
				switch matrix[i][j] {
				case 0:
					matrix[i][j] = 'u'
				case 'u':
					overlap++
					matrix[i][j] = 'x'
				default:
					matrix[i][j] = ' '
				}
			}
		}
	}

	uni := make(map[int]bool)
	for k, cl := range claims {
		offx, offy := coord(cl[0])
		x, y := coord(cl[1])
		var u []bool
		for i := offx; i < offx+x; i++ {
			for j := offy; j < offy+y; j++ {
				switch matrix[i][j] {
				case 'x':
					u = append(u, false)
				case 'u':
					u = append(u, true)
				default:
					u = append(u, false)
				}
			}
		}
		un := true
		for _, v := range u {
			if !v {
				un = false
			}
		}
		uni[k] = un
	}

	var ucl int
	for k, v := range uni {
		if v {
			ucl = k
		}
	}

	fmt.Printf("There is %d square inches of overlap.\n", overlap)
	fmt.Printf("The unique claim is %d.\n", ucl)
}

func conv(str []string) []int {
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

func coord(x []int) (int, int) {
	return x[0], x[1]
}
