package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	matrix := [][]float64{
		{4, 3, 1, 0},
		{-2, 2, 6, 1},
		{0, 5, 2, 3},
		{0, 1, 2, 7},
	}

	vector := []float64{29, 38, 48, 56}

	u, m, x, det := gauss(matrix, vector)

	fmt.Print("\nMatrix U:")

	for i := 0; i < len(matrix); i++ {
		fmt.Print("\n\t")

		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(strconv.FormatFloat(math.Round(u[i][j]*100)/100, 'f', -1, 64) + "\t")
		}
	}

	fmt.Print("\nMatrix M:")

	for i := 0; i < len(matrix); i++ {
		fmt.Print("\n\t")

		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(strconv.FormatFloat(math.Round(m[i][j]*100)/100, 'f', -1, 64) + "\t")
		}
	}

	fmt.Print("\n\nVector x: ")

	for i := 0; i < len(matrix); i++ {
		fmt.Printf("x[%d] = %s; ", i+1, strconv.FormatFloat(x[i], 'f', -1, 64))
	}

	fmt.Printf("\n\ndet A = %s", strconv.FormatFloat(det, 'f', -1, 64))

	revMatrix := createEmptyMatrix(len(matrix))

	for i := 0; i < len(matrix); i++ {
		currVector := make([]float64, len(matrix))

		for j := 0; j < len(matrix); j++ {
			currVector[j] = m[j][i]
		}

		_, _, currCol, _ := gauss(u, currVector)

		for j := 0; j < len(matrix); j++ {
			revMatrix[j][i] = currCol[j]
		}
	}

	fmt.Print("\n\nMatrix A^-1:")

	for i := 0; i < len(matrix); i++ {
		fmt.Print("\n\t")

		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(strconv.FormatFloat(math.Round(revMatrix[i][j]*100)/100, 'f', -1, 64) + "\t")
		}
	}
}
