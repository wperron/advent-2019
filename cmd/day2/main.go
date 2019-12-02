package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	handle, _ := os.Open("Intcode.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)

	for scanner.Scan() {
		intcode := ToIntSlice(strings.Split(scanner.Text(), ","))
		repaired := Repair(intcode)
		calculated := Intcode(repaired)
		fmt.Printf("Intcode at pos 0: %d\n", calculated[0])
	}
}

func ToIntSlice(x []string) ([]int) {
	y := make([] int, len(x))
	for i := range x {
		y[i], _ = strconv.Atoi(x[i])
	}
	return y
}

func Repair(intcode []int) []int {
	intcode[1] = 12
	intcode[2] = 2
	return intcode
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
