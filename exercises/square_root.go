// Author: [Hilson-Alex](https://github.com/Hilson-Alex).

package exercises

import (
	"fmt"
	"math"
)

var pow = math.Pow

// ErrNegativeSqrt is an error type for negative numbers,
// as asked in https://go.dev/tour/methods/20.
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("Square root not supported for negative numbers as ", float64(e))
}

// SquareRoot computes the approximate square root of a positive number
// or throws an error if the number is negative, as instructed in
// https://go.dev/tour/flowcontrol/8.
func SquareRoot(value float64) (float64, error) {
	if value < 0 {
		return 0, ErrNegativeSqrt(value)
	}
	var result = value / 2
	for math.Abs(pow(result, 2)-value) > 0.0000000000001 {
		result -= (pow(result, 2) - value) / (2 * result)
	}
	return result, nil
}
