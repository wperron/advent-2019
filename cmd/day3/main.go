package main

import (
	"bufio"
	"fmt"
	"math"
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
	handle, _ := os.Open("./debug.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)
	var paths [][]vector

	for scanner.Scan() {
		paths = append(paths, CalcPath(strings.Split(scanner.Text(), ",")))
	}

	for _, path := range paths {
		fmt.Println(len(path))
	}

	dist := Solve(paths[0], paths[1])
	fmt.Println("Distance: ", dist)
}

func Solve(red []vector, green []vector) int {
	var short int

	for _, v1 := range red {
		for _, v2 := range green {
			cross, intersects := CrossPoint(v1, v2)
			if intersects {
				dist := int(math.Abs(float64(cross.x)) + math.Abs(float64(cross.y)))
				if short == 0 && dist > 0 {
					short = dist
				} else if dist < short {
					fmt.Println("shortest: ", cross)
					short = dist
				}
			}
		}
	}

	return short
}

func CrossPoint(a vector, b vector) (intTuple, bool) {
	var cross intTuple
	xdiff := intTuple{a.from.x - a.to.x, b.from.x - b.to.x}
	ydiff := intTuple{a.from.y - a.to.y, b.from.y - b.to.y}

	det := func (a intTuple, b intTuple) int {
		return a.x * b.y - a.y * b.x
	}

	div := det(xdiff, ydiff)
	if div == 0 {
		return cross, false
	}

	d := intTuple{det(a.from, a.to), det(b.from, b.to)}
	x := det(d, xdiff) / div
	y := det(d, ydiff) / div

	return intTuple{x, y}, true
}

// func Intersects(a vector, b vector) bool {
// 	if (a.from.x == a.to.x && b.from.x == b.to.x) || (a.from.y == a.to.y && b.from.y == b.to.y) {
// 		return false // assume closest intersection is not on colinear vectors,
// 	} else if (math.Min(float64(a.from.x), float64(a.to.x)) < float64(b.from.x)) && 
// 		(float64(b.from.x) < math.Max(float64(a.from.x), float64(a.to.x))) &&
// 		(math.Min(float64(a.from.y), float64(a.to.y)) < float64(b.from.y)) &&
// 		(float64(b.from.y) < math.Max(float64(a.from.y), float64(a.to.y))) {
// 		return true
// 	} else {
// 		return false
// 	}
// }

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
		dest = intTuple{from.x, from.y + amt}
	case 'R':
		dest = intTuple{from.x + amt, from.y}
	case 'D':
		dest = intTuple{from.x, from.x - amt}
	case 'L':
		dest = intTuple{from.x - amt, from.y}
	}
	return dest
}
