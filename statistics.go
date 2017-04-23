// statistics package that provides some useful statistics functions
package statistics

import (
	"errors"
	"math"
)

// Calculate the mean of a slice (x)
func Mean(x []float64) (mean float64) {
	// Avoid divide by zero
	if len(x) == 0 {
		return
	}
	// Sum all values from the x slice
	for index := 0; index < len(x); index++ {
		mean += x[index]
	}
	// Calculate the mean by dividing by the number of elements in x
	mean /= float64(len(x))
	return
}

// Calculate the standard deviation of a sample (n - 1)
func StandardDeviation(x []float64) (sd float64, err error) {
	// Check if the slice is not empty
	if len(x) <= 1 {
		err = errors.New("The slice has less than 1 element")
		return
	}

	// Calculate the x mean
	xMean := Mean(x)

	// Calculate the numerator
	for index := 0; index < len(x); index++ {
		sd += math.Pow((x[index] - xMean), 2)
	}

	// Divide by (n - 1)
	sd /= float64(len(x) - 1)

	// Calculate the square root
	sd = math.Sqrt(sd)
	return
}

// Calculate the variance
func Variance(x []float64) (variance float64, err error) {
	// Calculate the standard deviation
	sd, err := StandardDeviation(x)

	// If there was some error in the standard deviation function
	if err != nil {
		return
	}

	// Calculate the variance
	variance = math.Pow(sd, 2)
	return
}

// Calculate the covariance by the formula COV(X,Y) = sum((Xi-Xm)(Yi-Ym)) / (n-1)
func Covariance(x []float64, y []float64) (result float64, err error) {
	// Check if we have two valid slices
	if len(x) != len(y) {
		err = errors.New("The two slices have different sizes")
		return
	}
	if len(x) == 0 || len(y) == 0 {
		err = errors.New("The slices are empty")
		return
	}

	// Calculate the mean of X and Y
	xMean := Mean(x)
	yMean := Mean(x)

	// Calculate the numerator sum((Xi-Xm)(Yi-Ym))
	for index := 0; index < len(x); index++ {
		result += (x[index] - xMean) * (y[index] - yMean)
	}

	// Divide by the denominator (n-1)
	result /= float64((len(x) - 1))
	return
}

// Calculate the covariance matrix
func CovarianceMatrix(matrix [][]float64) (result [][]float64, err error) {
	// Check if the matrix has at least one row
	if len(matrix) == 0 {
		err = errors.New("Empty matrix")
		return
	}

	// Check if the first row has at least one column
	if len(matrix[0]) == 0 {
		err = errors.New("Empty matrix")
		return
	}

	// For each row in the matrix calculate the covariance of two dimensions (this is the first dimension)
	for row := 0; row < len(matrix); row++ {

		// Create a new row in the result covariance matrix
		result = append(result, []float64{})

		// For each row calculate the covariance (this is the second dimension)
		for tempRow := 0; tempRow < len(matrix); tempRow++ {

			// Calculate the covariance of the two rows/dimensions
			covResult, covError := Covariance(matrix[row], matrix[tempRow])

			// If could not calculate the covariance of the two rows, return an error
			if covError != nil {
				result = nil // We need to set the result as nil
				err = errors.New("Error trying to calculate the covariance matrix")
				return
			}

			// Append the result to the result matrix based on the current row (first dimension)
			result[row] = append(result[row], covResult)
		}
	}

	return
}
