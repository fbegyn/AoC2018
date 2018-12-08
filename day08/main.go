package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type states int

const (
	numChildren = iota
	numMetadata
	metadata
)

func (s states) String() string {
	names := [...]string{
		"numChildren",
		"numMetadata",
		"metadata",
	}
	return names[s]
}

type node struct {
	state      states
	childNodes int
	metas      int
	value      int
	children   []*node
}

func newNode() *node {
	return &node{numChildren, 0, 0, 0, nil}
}

func main() {
	f, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.Trim(string(f), "\n"), " ")

	curNode := newNode()
	root := curNode
	var stack []*node

	sum := 0
	for _, in := range input {
		num, err := strconv.Atoi(in)
		if err != nil {
			panic(err)
		}

		switch curNode.state {
		case numChildren:
			curNode.childNodes = num
			curNode.state = numMetadata
		case numMetadata:
			curNode.metas = num
			curNode.state = metadata

			if curNode.childNodes > 0 {
				stack = append(stack, curNode)
				nextNode := newNode()
				curNode.children = append(curNode.children, nextNode)
				curNode = nextNode
			}
		case metadata:
			sum += num
			if len(curNode.children) == 0 {
				curNode.value += num
			} else {
				offset := num - 1
				if offset >= 0 && offset < len(curNode.children) {
					v := curNode.children[offset].value
					curNode.value += v
				}
			}
			curNode.metas--
			if curNode.metas == 0 {
				if len(stack) > 0 {
					curNode = stack[len(stack)-1]
					curNode.childNodes--
					if curNode.childNodes > 0 {
						nextNode := newNode()
						curNode.children = append(curNode.children, nextNode)
						curNode = nextNode
					} else {
						stack = stack[:len(stack)-1]
					}
				} else {
					curNode = newNode()
					stack = nil
				}
			}
		}
	}
	log.Printf("Sum of all metadata: %d\n", sum)
	log.Printf("Value of the root node: %d\n", root.value)
}
