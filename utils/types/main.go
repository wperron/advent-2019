package types

import (
	"strconv"
)

type IntTuple struct {
	X, Y int
}

type Vector struct {
	From, To IntTuple
}

func MakeRange(min int, max int) []int {
	res := make([]int, max-min)
	for i := range res {
		res[i] = min + i
	}
	return res
}

func Permutations(a []int, b []int) []IntTuple {
	var permuts []IntTuple
	for _, i := range a {
		for _, j := range b {
			permuts = append(permuts, IntTuple{i, j})
		}
	}
	return permuts
}

func ToIntSlice(x []string) []int {
	y := make([]int, len(x))
	for i := range x {
		y[i], _ = strconv.Atoi(x[i])
	}
	return y
}