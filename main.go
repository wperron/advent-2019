package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"math"
)

func main() {
	handle, _ := os.Open("module-weights.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)

	var total int
	for scanner.Scan() {
		module, _ := strconv.Atoi(scanner.Text())
		mass := CalcFuel(module)
		total += mass

		fmt.Printf("Module: %d\tFuel: %d\n", module, mass)
	}
	fmt.Printf("Total: %d", total)
	os.Exit(0)
}

func CalcFuel(mass int) int {
	fuel := int(math.Floor(float64(mass)/3))-2
	if fuel > 0 {
		fuel += CalcFuel(fuel)
	}
	return int(math.Max(float64(fuel), 0.0))
}