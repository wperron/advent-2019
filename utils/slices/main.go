package slices

func Occurences(s []int, symbol int) (total int) {
	for _, v := range s {
		if v == symbol {
			total++
		}
	}
	return
}