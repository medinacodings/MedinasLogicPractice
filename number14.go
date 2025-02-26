package Logic3

import "github.com/aronipurwanto/go-basic-logic/utils"

func Number14(n int) (result [][]int) {
	// Create a 2D slice
	result = utils.CreateSlice(n)

	// Fill the slice
	for col := 0; col < n; col++ {
		num := 1
		for row := 0; row < n; row++ {
			if col%2 == 0 {
				// For even columns, fill normally
				result[row][col] = num
			} else if row-col >= 0 { // Ensure the index is valid
				// For odd columns, fill "reverse diagonal" style
				result[row-col][col] = num
			}
			num += 2
		}
	}
	return result
}
