package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	handle, _ := os.Open("./range.txt")
	defer handle.Close()
	scanner := bufio.NewScanner(handle)

	for scanner.Scan() {
		rangeDef := strings.Split(scanner.Text(), "-")
		min, _ := strconv.Atoi(rangeDef[0])
		max, _ := strconv.Atoi(rangeDef[1])
		var possible []string

		for i := min; i <= max; i++ {
			trial := strconv.Itoa(i)
			if IsValid(trial) {
				possible = append(possible, trial)
			}
		}
		fmt.Println("possible combinations: ", len(possible))
		break
	}
}

func IsValid(pwd string) bool {
	seq := strings.Split(pwd, "")
	inOrder := sort.StringsAreSorted(seq)

	count := 1
	var last string
	var hasSeq bool
	for _, ch := range append(seq, "") {
		if ch == last {
			count++
		} else {
			if count == 2 {
				hasSeq = true
			}
			count = 1
		}
		last = ch
	}

	return inOrder && hasSeq
}
