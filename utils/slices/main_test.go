package slices

import (
	"testing"
)

func TestOccurences(t *testing.T) {
	a := []int{0, 1, 2, 3, 4}
	b := []int{0, 0, 0, 1}
	var actual int

	actual = Occurences(a, 0)
	if actual != 1 {
		t.Errorf("Error: expected [ %d ] got [ %d ]", 1, actual)
	}

	actual = Occurences(b, 0)
	if actual != 3 {
		t.Errorf("Error: expected [ %d ] got [ %d ]", 3, actual)
	}

	actual = Occurences(b, 5)
	if actual != 0 {
		t.Errorf("Error: expected [ %d ] got [ %d ]", 0, actual)
	}
}