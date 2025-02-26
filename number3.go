package Logic3

func Number3(n int) (result [][]int) {
	result = make([][]int, n) // Create a 2D slice with `n` rows
	value := 2                // Starting value for the first row

	for i := 0; i < n; i++ {
		result[i] = make([]int, n-i) // Each row has `n-i` elements

		for j := 0; j < len(result[i]); j++ {
			result[i][j] = value

			// Alternate between adding and subtracting
			if i%2 == 0 {
				value += 3 // Increment the value by 3 for even rows
			} else {
				value -= 3 // Decrement the value by 3 for odd rows
			}
		}
	}
	return result
}
