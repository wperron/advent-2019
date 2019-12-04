package main

import (
	"testing"
)

type pair struct {
	in string
	out bool
}

func TestIsValid(t *testing.T) {
	cases := []pair{
		pair{"111111", false},
		pair{"234567", false},
		pair{"112233", true},
		pair{"111122", true},
		pair{"223333", true},
		pair{"223450", false},
		pair{"123789", false},
		pair{"123444", false},
	}

	for _, this := range cases {
		actual := IsValid(this.in)
		if actual != this.out {
			t.Error("Wrong assertion for: ", this, actual)
		}
	}
}