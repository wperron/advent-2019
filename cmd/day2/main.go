package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var TARGET int = 19690720

type intTuple struct {
	a, b int
}

func main() {
	handle, _ := os.Open("./Intcode.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)
	c := make(chan intTuple)

	nouns := MakeRange(0, 100)
	verbs := MakeRange(0, 100)
	permuts := Permutations(nouns, verbs)

	for scanner.Scan() {
		intcode := ToIntSlice(strings.Split(scanner.Text(), ","))

		for _, p := range permuts {
			go func(intcode []int, in intTuple) {
				copied := make([]int, len(intcode))
				copy(copied, intcode)

				copied[1] = in.a
				copied[2] = in.b

				res := Intcode(copied)[0]

				if res == TARGET {
					c <- in
					close(c)
				}
			}(intcode, p)
		}

		answer := <-c

		fmt.Println("Answer is: ", answer)
	}
}

func MakeRange(min int, max int) []int {
	res := make([]int, max-min)
	for i := range res {
		res[i] = min + i
	}
	return res
}

func Permutations(a []int, b []int) []intTuple {
	var permuts []intTuple
	for _, i := range a {
		for _, j := range b {
			permuts = append(permuts, intTuple{i, j})
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

func Intcode(intcode []int) []int {
	length := len(intcode)
	for i := 0; i < length; i += 4 {
		if intcode[i] == 99 {
			break
		}

		opcode := intcode[i]
		operands := intcode[i+1 : i+3]
		target := intcode[i+3]

		switch opcode {
		case 1:
			intcode[target] = intcode[operands[0]] + intcode[operands[1]]
		case 2:
			intcode[target] = intcode[operands[0]] * intcode[operands[1]]
		default:
			panic("The opcode is unkown")
		}
	}
	return intcode
}
