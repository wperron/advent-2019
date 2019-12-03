package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type intTuple struct {
	x, y int
}

type vector struct {
	from, to intTuple
}

func main() {
	fmt.Println("Hello World!")
	handle, _ := os.Open("./coordinates.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)
	var paths [][]vector

	for scanner.Scan() {
		paths = append(paths, CalcPath(strings.Split(scanner.Text(), ",")))
	}

	for _, path := range paths {
		fmt.Println(len(path))
	}
}

func CalcPath(instructions []string) []vector {
	path := make([]vector, len(instructions))
	start := intTuple{0, 0}
	last := start
	for i, move := range instructions {
		next := MoveToCoordinate(last, move)
		path[i] = vector{last, next}
		last = next
	}
	return path
}

func MoveToCoordinate(from intTuple, move string) intTuple {
	var dest intTuple
	dir := move[0]
	amt, _ := strconv.Atoi(move[1:])
	switch dir {
	case 'U':
		dest = intTuple{from.x, from.y+amt}
	case 'R':
		dest = intTuple{from.x+amt, from.y}
	case 'D':
		dest = intTuple{from.x, from.x-amt}
	case 'L':
		dest = intTuple{from.x-amt, from.y}
	}
	return dest
}