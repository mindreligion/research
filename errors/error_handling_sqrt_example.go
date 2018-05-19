package main

import (
	_ "errors"
	"fmt"
	"math"
)

type NegativeSqrtError float64

func (err NegativeSqrtError) Error() string {
	return fmt.Sprintf(`Argument should be not negative, %g is given`, float64(err))
}

func getSqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, NegativeSqrtError(f)
	}
	return math.Sqrt(f), nil
}

func main() {
	for i := -2; i < 3; i++ {
		sqrt, err := getSqrt(float64(i))
		if err != nil {
			println(err.Error())
			continue
		}
		println(sqrt)
	}
}
