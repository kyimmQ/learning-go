package basics

import (
	"fmt"
	"math"
)

var Global int = 1234
var AnoterGlobal int = -5678

const PI float64 = 3.14159

func Variable() {
	var j int
	i := Global + AnoterGlobal

	j = Global

	k := math.Abs(float64(AnoterGlobal))

	fmt.Printf("i = %d, j = %d, k = %f\n", i, j, k)
	fmt.Printf("PI = %f\n", PI)
}
