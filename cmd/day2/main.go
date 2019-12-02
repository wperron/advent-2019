package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	handle, _ := os.Open("./Intcode.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)

	for scanner.Scan() {
		intcode := ToIntSlice(strings.Split(scanner.Text(), ","))

		res, err := Iterate(intcode)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("Intcode at pos 0: %d\n", res[0])
	}
}

func ToIntSlice(x []string) ([]int) {
	y := make([] int, len(x))
	for i := range x {
		y[i], _ = strconv.Atoi(x[i])
	}
	return y
}

func Repair(intcode []int, noun int, verb int) []int {
	intcode[1] = noun
	intcode[2] = verb
	return intcode
}

func Iterate(intcode []int) ([]int, error) {
	for i := 0; 1 < 100; i++ {
		for j := 0; j < 100; j++ {
			copied := make([]int, len(intcode))
			copy(copied, intcode)
			repaired := Repair(copied, i, j)

			calculated := Intcode(repaired)
			if calculated[0] == 19690720 {
				fmt.Printf("Noun: %d\tVerb: %d\n", i, j)
				return calculated, nil
			}
		}
	}
	return nil, errors.New("Permutation not found.")
}

func Intcode(intcode []int) []int {
	length := len(intcode)
	for i := 0; i < length; i += 4 {
		if intcode[i] == 99 { break }

		opcode := intcode[i]
		operands := intcode[i+1:i+3]
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
