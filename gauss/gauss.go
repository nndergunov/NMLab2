package main

import (
	"math"
)

func gauss(A [][]float64, b []float64) ([][]float64, [][]float64, []float64, float64) {
	u := copyMatrix(A)
	vector := copyVector(b)
	size := len(u)
	largestEl := make([]float64, size)
	swaps := 0

	m := createDiagonalMatrix(size)

	for i := 0; i < size; i++ {
		// choose the largest element
		maxPos := i

		for j := i; j < size; j++ {
			if abs(u[j][i]) > abs(u[maxPos][i]) {
				maxPos = j
			}
		}

		// check if solvable
		if u[maxPos][i] == 0 {
			panic("max index = 0")
		}

		largestEl[i] = u[maxPos][i]

		// swap i-th and j-th rows
		if maxPos != i {
			swaps++

			p := createDiagonalMatrix(size)

			u[i], u[maxPos] = u[maxPos], u[i]
			vector[i], vector[maxPos] = vector[maxPos], vector[i]
			p[i], p[maxPos] = p[maxPos], p[i]

			m = multiplyMatrices(p, m)
		}

		currM := createDiagonalMatrix(size)

		currM[i][i] = 1 / u[i][i]
		for j := i + 1; j < size; j++ {
			currM[j][i] = -u[j][i] / u[i][i]
		}

		m = multiplyMatrices(currM, m)

		// divide i-th row by u[i][i]
		for j := i + 1; j < size; j++ {
			u[i][j] /= u[i][i]
		}

		vector[i] /= u[i][i]

		u[i][i] = 1

		// Subtract i-th row from j-th row (j = 1, 2, ..., size and j != i)
		for j := i; j < size; j++ {
			if j != i {
				for k := i + 1; k < size; k++ {
					u[j][k] -= u[i][k] * u[j][i]
				}

				vector[j] -= vector[i] * u[j][i]

				u[j][i] = 0
			}
		}
	}

	answers := make([]float64, size)

	for i := size - 1; i >= 0; i-- {
		res := vector[i]

		for j := i + 1; j < size; j++ {
			res -= u[i][j] * answers[j]
		}

		answers[i] = res / u[i][i]
	}

	return u, m, answers, findDet(largestEl, swaps)
}

func findDet(largestEl []float64, n int) float64 {
	det := math.Pow(-1, float64(n))

	for i := 0; i < len(largestEl); i++ {
		det *= largestEl[i]
	}

	return math.Round(det)
}

func multiplyMatrices(matrix1, matrix2 [][]float64) [][]float64 {
	m1 := copyMatrix(matrix1)
	m2 := copyMatrix(matrix2)
	res := createEmptyMatrix(len(matrix1))

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1); j++ {
			res[i][j] = 0

			for k := 0; k < len(matrix1); k++ {
				el := m1[i][k] * m2[k][j]

				res[i][j] += el
			}
		}
	}

	return res
}

func abs(n float64) float64 {
	return math.Abs(n)
}

func copyVector(vector []float64) []float64 {
	newVector := make([]float64, len(vector))

	copy(newVector, vector)

	return newVector
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

func createDiagonalMatrix(size int) [][]float64 {
	m := createEmptyMatrix(size)
	for i := 0; i < size; i++ {
		m[i][i] = 1
	}

	return m
}

func createEmptyMatrix(size int) [][]float64 {
	m := make([][]float64, size)

	for i := 0; i < size; i++ {
		row := make([]float64, size)

		m[i] = row
	}

	return m
}
