package Logic3

import "github.com/aronipurwanto/go-basic-logic/utils"

func Number9(n int) (result [][]int) {
	result = utils.CreateSlice(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i][j] = i + j + 1
		}
	}
	return result
}
