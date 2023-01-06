// Author: [Hilson-Alex](https://github.com/Hilson-Alex).

package exercises

// Fibonacci is a closure that returns a function for
// calculate the next number of the fibonacci sequence
// as instructed in https://go.dev/tour/moretypes/26.
func Fibonacci() func() int {
	var current, next = 0, 1
	return func() int {
		var aux = current
		current, next = next, current+next
		return aux
	}
}
