package main

func tridiagonalMatrixAlgorithm(matrix [][]float64, vector []float64) []float64 {
	n := len(matrix)

	if n <= 2 {
		return vector
	}

	matrix[0][1] /= matrix[0][0]
	vector[0] /= matrix[0][0]

	for i := 1; i < n-1; i++ {
		matrix[i][i+1] /= matrix[i][i] - matrix[i][i-1]*matrix[i-1][i]
	}

	for i := 1; i < n; i++ {
		vector[i] = (vector[i] - matrix[i][i-1]*vector[i-1]) / (matrix[i][i] - matrix[i][i-1]*matrix[i-1][i])
	}

	x := make([]float64, n)

	x[n-1] = vector[n-1]

	for i := n - 2; i >= 0; i-- {
		x[i] = vector[i] - matrix[i][i+1]*x[i+1]
	}

	return x
}
