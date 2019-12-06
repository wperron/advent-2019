package wires

import (
	"testing"
)

func TestCrossPoint(t *testing.T) {
	var v1, v2 vector
	var expects, actual intTuple
	var intersects bool
	
	v1 = vector{intTuple{0, 0}, intTuple{10, 0}}
	v2 = vector{intTuple{5, -5}, intTuple{5, 5}}
	expects = intTuple{5, 0}

	actual, intersects = CrossPoint(v1, v2)
	if !intersects {
		t.Error("Error crosspoint", v1, v2, actual, intersects)
	}
	if actual != expects {
		t.Error("Error crosspoint", v1, v2, actual, intersects)
	}

	v1 = vector{intTuple{0, 0}, intTuple{10, 0}}
	v2 = vector{intTuple{0, 10}, intTuple{10, 10}}

	actual, intersects = CrossPoint(v1, v2)
	if intersects {
		t.Error("Error crosspoint", v1, v2, actual, intersects)
	}
}