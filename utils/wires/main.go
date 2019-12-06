package wires

import (
	"fmt"
	"github.com/wperron/advent/utils/types"
	"math"
	"strconv"
)

func Solve(red []types.Vector, green []types.Vector) int {
	var short int

	for _, v1 := range red {
		for _, v2 := range green {
			cross, intersects := CrossPoint(v1, v2)
			if intersects {
				dist := int(math.Abs(float64(cross.X)) + math.Abs(float64(cross.Y)))
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

// func CrossPoint(a types.Vector, b types.Vector) (types.IntTuple, bool) {
// 	var cross types.IntTuple
// 	xdiff := types.IntTuple{a.from.x - a.to.x, b.from.x - b.to.x}
// 	ydiff := types.IntTuple{a.from.y - a.to.y, b.from.y - b.to.y}

// 	det := func (a types.IntTuple, b types.IntTuple) int {
// 		return a.x * b.y - a.y * b.x
// 	}

// 	div := det(xdiff, ydiff)
// 	if div == 0 {
// 		return cross, false
// 	}

// 	d := types.IntTuple{det(a.from, a.to), det(b.from, b.to)}
// 	x := det(d, xdiff) / div
// 	y := det(d, ydiff) / div

// 	return types.IntTuple{x, y}, true
// }

func CalcPath(instructions []string) []types.Vector {
	path := make([]types.Vector, len(instructions))
	start := types.IntTuple{0, 0}
	last := start
	for i, move := range instructions {
		next := MoveToCoordinate(last, move)
		path[i] = types.Vector{last, next}
		last = next
	}
	return path
}

// func MoveToCoordinate(from types.IntTuple, move string) types.IntTuple {
// 	var dest types.IntTuple
// 	dir := move[0]
// 	amt, _ := strconv.Atoi(move[1:])
// 	switch dir {
// 	case 'U':
// 		dest = types.IntTuple{from.x, from.y + amt}
// 	case 'R':
// 		dest = types.IntTuple{from.x + amt, from.y}
// 	case 'D':
// 		dest = types.IntTuple{from.x, from.x - amt}
// 	case 'L':
// 		dest = types.IntTuple{from.x - amt, from.y}
// 	}
// 	return dest
// }
