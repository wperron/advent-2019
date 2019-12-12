package main

import (
	"bufio"
	"fmt"
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

	/*
		0 => black
		1 => white
		2 => transparent
	*/

	var curr []int
	var layers [][]int
	for len(pixels) > 0 {
		curr, pixels = pixels[:width*height], pixels[width*height:]
		layers = append(layers, curr)
	}

	fmt.Println("first layer: ", layers[0])

	rendered := make([]int, width*height)
	for i := 0; i < width*height; i++ {
		for _, l := range layers {
			if l[i] == 2 {
				continue
			} else {
				rendered[i] = l[i]
				break
			}
		}
	}

	// convert layers to more readable characters
	renderedText := []string{}
	for _, v := range rendered {
		if v == 0 {
			renderedText = append(renderedText, " ")
		} else {
			renderedText = append(renderedText, "#")
		}
	}

	fmt.Println("Final render: ")
	var line []string
	for i := 0; i < height; i++ {
		line, renderedText = renderedText[:width], renderedText[width:]
		fmt.Println(strings.Join(line, ""))
	}
}
