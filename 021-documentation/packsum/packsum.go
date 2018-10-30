// Package packsum provides math functions
package packsum

// Sum adds unlimited number of values of type int
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
