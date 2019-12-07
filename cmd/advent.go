package main

import (
	"bufio"
	"fmt"
	"github.com/wperron/advent/utils/graph"
	"strings"
	"os"
)

var TARGET int = 19690720

func main() {
	handle, _ := os.Open("../resources/orbits.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)

	gr := graph.Graph{}
	for scanner.Scan() {
		orbit := strings.Split(scanner.Text(), ")")
		for _, n := range orbit {
			graph.AddNode(&gr, graph.Node{n})
		}
		graph.AddEdge(&gr, graph.Node{orbit[0]}, graph.Node{orbit[1]})
	}

	fmt.Println("Using graph: ", gr)
	var total int
	for _, node := range graph.Nodes(gr) {
		total += graph.Traverse(gr, node)
	}
	fmt.Println(total)
}