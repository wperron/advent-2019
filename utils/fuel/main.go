package fuel

import (
	"math"
)

func CalcFuel(mass int) int {
	fuel := int(math.Floor(float64(mass)/3)) - 2
	if fuel > 0 {
		fuel += CalcFuel(fuel)
	}
	return int(math.Max(float64(fuel), 0.0))
}
