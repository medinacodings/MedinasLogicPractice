package Logic3

import go_print_slice "github.com/Florenzee/go-print-slice"

func Number13(n int) (result [][]int) {
	result = go_print_slice.CreateSlice(n)

	mid := n / 2
	values := 1

	// Build the diamond shape
	for i := 0; i <= mid; i++ {
		for j := 0; j <= i; j++ {
			if (i%2 == 0 && j%2 == 0) || (i%2 == 1 && j%2 == 1) {
				// Fill left and right side symmetrically
				result[mid-j][i] = values
				result[mid+j][i] = values
				result[mid-j][n-i-1] = values
				result[mid+j][n-i-1] = values
			}
		}
		values += 2 // Increment the value by 2
	}
	return result
}
