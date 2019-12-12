package main

import (
	"bufio"
	"fmt"
	"github.com/wperron/advent/utils/slices"
	"os"
	"strconv"
	"strings"
)

var TARGET int = 19690720

func main() {
	handle, _ := os.Open("../resources/pwd-layers.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)

	var pixels []int
	width := 25
	height := 6
	for scanner.Scan() {
		tmp := strings.Split(scanner.Text(), "")
		for _, p := range tmp {
			pInt, _ := strconv.Atoi(p)
			pixels = append(pixels, pInt)
		}
	}

	fmt.Println("Pixels: ", len(pixels))

	var curr []int
	var fewest, ones, twos int
	for len(pixels) > 0 {
		curr, pixels = pixels[:width*height], pixels[width*height:]
		zeros := slices.Occurences(curr, 0)

		if zeros < fewest || fewest == 0 {
			fewest = zeros

			ones = slices.Occurences(curr, 1)
			twos = slices.Occurences(curr, 2)
		}
	}

	fmt.Println("Ones: ", ones)
	fmt.Println("Twos: ", twos)
	fmt.Println("Answer: ", ones * twos)
}