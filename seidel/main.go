package main

import (
	"fmt"
	"strconv"
)

func main() {
	matrix := [][]float64{
		{5, 1, 1, 0},
		{1, 2, 0, 0},
		{1, 0, 4, 2},
		{0, 0, 2, 3},
	}

	vector := []float64{10, 5, 21, 18}

	x := seidel(matrix, vector)

	fmt.Println("vector x: ")

	for i := 0; i < len(matrix); i++ {
		fmt.Printf("\tx[%d] = %s\n", i+1, strconv.FormatFloat(x[i], 'f', -1, 64))
	}
}
