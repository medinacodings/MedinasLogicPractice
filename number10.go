package Logic3

import (
	go_print_slice "github.com/Florenzee/go-print-slice"
)

func Number10(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)

	mid := (n - 1) / 2
	for i := mid; i >= 0; i-- {
		value := 1 + (2 * (mid - i))
		for j := mid; j >= mid-i; j-- {
			result[i][j] = value
			result[i][n-1-j] = value
			result[n-1-i][j] = value
			result[n-1-i][n-1-j] = value
			value += 2
		}
	}
	return result
}
