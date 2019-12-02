package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	handle, _ := os.Open("module-weights.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)

	var total int
	masses := make(chan int)
	for scanner.Scan() {
		module, _ := strconv.Atoi(scanner.Text())
		go func(i int) {
			masses <- CalcFuel(i)
		}(module)

		total += <- masses
	}
	fmt.Printf("Total: %d", total)
	os.Exit(0)
}

func CalcFuel(mass int) int {
	fuel := int(math.Floor(float64(mass)/3)) - 2
	if fuel > 0 {
		fuel += CalcFuel(fuel)
	}
	return int(math.Max(float64(fuel), 0.0))
}
