package main

import (
	"bufio"
	"fmt"
	"github.com/wperron/advent/utils/intcode"
	"github.com/wperron/advent/utils/types"
	"strings"
	"os"
)

var TARGET int = 19690720

func main() {
	handle, _ := os.Open("../resources/Intcode.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)
	c := make(chan types.IntTuple)

	nouns := types.MakeRange(0, 100)
	verbs := types.MakeRange(0, 100)
	permuts := types.Permutations(nouns, verbs)

	for scanner.Scan() {
		seq := types.ToIntSlice(strings.Split(scanner.Text(), ","))

		for _, p := range permuts {
			go func(seq []int, in types.IntTuple) {
				copied := make([]int, len(seq))
				copy(copied, seq)

				copied[1] = in.X
				copied[2] = in.Y

				res := intcode.Intcode(copied)[0]

				if res == TARGET {
					c <- in
					close(c)
				}
			}(seq, p)
		}

		answer := <-c

		fmt.Println("Answer is: ", answer)
	}
}