package main

import (
	"fmt"
	"math"
)

func seidel(a [][]float64, b []float64) []float64 {
	n := len(a)

	// checking if a is symmetrical
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if a[i][j] != a[j][i] {
				panic("impossible to use seidel method")
			}
		}
	}

	// checking if determinants are positive
	for i := 1; i <= n; i++ {
		subMatrix := make([][]float64, i)

		for j := 0; j < i; j++ {
			row := make([]float64, i)

			for k := 0; k < i; k++ {
				row[k] = a[j][k]
			}

			subMatrix[j] = row
		}

		if findDet(subMatrix) <= 0 {
			panic("det is zero or negative")
		}
	}

	xPrev := make([]float64, n)
	x := make([]float64, n)
	y := make([]float64, n)
	eps := 0.001
	iter := 0

	for counter := 1; ; counter++ {
		iter++

		copy(xPrev, x)

		for i := 0; i < n; i++ {
			y[i] = b[i] / a[i][i]

			for j := 0; j < n; j++ {
				if j == i {
					continue
				}

				y[i] -= (a[i][j] / a[i][i]) * x[j]

				x[i] = y[i]
			}

			x[i] = y[i]
		}

		done := true

		for i := 0; i < n; i++ {
			if math.Abs(xPrev[i]-x[i]) > eps {
				done = false
			}
		}

		if done {
			fmt.Println("took", counter, "iterations")

			break
		}
	}

	for i := 0; i < n; i++ {
		x[i] = math.Round(x[i])
	}

	return x
}

func findDet(m [][]float64) float64 {
	matrix := copyMatrix(m)

	n := len(matrix)

	if n == 1 {
		return m[0][0]
	}

	if n == 2 {
		return (matrix[0][0] * matrix[1][1]) - (matrix[1][0] * matrix[0][1])
	}

	submatrix := createEmptyMatrix(n - 1)

	var det float64

	for x := 0; x < n; x++ {
		subi := 0

		for i := 1; i < n; i++ {
			subj := 0

			for j := 0; j < n; j++ {
				if j == x {
					continue
				}

				submatrix[subi][subj] = matrix[i][j]
				subj++
			}

			subi++
		}

		det += pow(-1, x) * matrix[0][x] * findDet(submatrix)
	}

	return det
}

func createEmptyMatrix(size int) [][]float64 {
	m := make([][]float64, size)

	for i := 0; i < size; i++ {
		row := make([]float64, size)

		m[i] = row
	}

	return m
}

func copyMatrix(matrix [][]float64) [][]float64 {
	newMatrix := make([][]float64, len(matrix))

	for i := 0; i < len(matrix); i++ {
		row := make([]float64, len(matrix[i]))

		copy(row, matrix[i])

		newMatrix[i] = row
	}

	return newMatrix
}

func pow(a, b int) float64 {
	return math.Pow(float64(a), float64(b))
}
