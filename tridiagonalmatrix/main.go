package main

import (
	"fmt"
	"strconv"
)

func main() {
	matrix := [][]float64{
		{1, 2, 0},
		{2, 2, 4},
		{0, 3, 3},
	}

	vector := []float64{5, 22, 18}

	x := tridiagonalMatrixAlgorithm(matrix, vector)

	fmt.Println("vector x: ")

	for i := 0; i < len(matrix); i++ {
		fmt.Printf("\tx[%d] = %s\n", i+1, strconv.FormatFloat(x[i], 'f', -1, 64))
	}
}
