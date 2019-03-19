package floyd

// matrix MUST be N * N array, and element of matrix MUST
// greater than 0
func Floyd(matrix [][]int) {
	for k := 0; k < len(matrix); k++ {
		for i := 0; i < len(matrix[k]); i++ {
			for j := 0; j < len(matrix[k]); j++ {
				if (matrix[i][k] + matrix[k][j]) < matrix[i][j] {
					matrix[i][j] = matrix[i][k] + matrix[k][j]
				}
			}
		}
	}
}
