package floyd

import (
	"fmt"
	"testing"
)

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%5d\t", matrix[i][j])
		}
		fmt.Println()
	}
}

func TestFloyd(t *testing.T) {
	const Infinity = 10000
	matrix := [][]int{
		{0, 12, Infinity, Infinity, Infinity, 16, 14},
		{12, 0, 10, Infinity, Infinity, 7, Infinity},
		{Infinity, 10, 0, 3, 5, 6, Infinity},
		{Infinity, Infinity, 3, 0, 4, Infinity, Infinity},
		{Infinity, Infinity, 5, 4, 0, 2, 8},
		{16, 7, 6, Infinity, 2, 0, 9},
		{14, Infinity, Infinity, Infinity, 8, 9, 0},
	}
	fmt.Println("before:")
	printMatrix(matrix)
	Floyd(matrix)
	fmt.Println("after:")
	printMatrix(matrix)
}
