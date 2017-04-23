
package statistics

import (
	"testing"
)

// Compare two float values using a predetermined epsilon
func floatEquals(a, b float64) bool {
	var epsilon float64 = 0.00000001
	if (a-b) < epsilon && (b-a) < epsilon {
		return true
	}
	return false
}

// Compare two matrices
func matrixEquals(a, b [][]float64) bool {
  if len(a) != len(b) {
    return false
  }

	for row := 0; row < len(a); row++ {
    if len(a[row]) != len(b[row]) {
      return false
    }

	  for col := 0; col < len(a[row]); col++ {
      if !floatEquals(a[row][col], b[row][col]) {
        return false
      }
	  }
	}

  return true
}

// Test the mean function
func TestMean(t *testing.T) {
  // Table tests
	var tTests = []struct {
		x    []float64
		mean float64
	}{
		{[]float64{}, 0},
  	{[]float64{0}, 0},
		{[]float64{10, 30}, 20},
		{[]float64{10, 10, 10}, 10},
		{[]float64{10, 20, 30, 40}, 25},
  }

	// Test with all values in the table
	for _, pair := range tTests {
		mean := Mean(pair.x)

		if !floatEquals(mean, pair.mean) {
			t.Error("Expected: ", pair.mean, "Received: ", mean)
		}
	}
}

// Test the standard deviation function
func TestStandardDeviation(t *testing.T) {

	_, err := StandardDeviation([]float64{})

  if err == nil {
    t.Error("Expected error different from nil")
  }

	_, err = StandardDeviation([]float64{1})

  if err == nil {
    t.Error("Expected error different from nil")
  }

  // Table tests
	var tTests = []struct {
		x  []float64
		sd float64
	}{
  	{[]float64{0, 50}, 35.35533905},
		{[]float64{10, 10, 10}, 0},
		{[]float64{10, 20, 30}, 10},
		{[]float64{0, 10, 15, 20, 35}, 12.9421791055},
  }

	// Test with all values in the table
	for _, pair := range tTests {
		sd, err := StandardDeviation(pair.x)

    if err != nil {
			t.Error("Expected: nil Received: ", err)
    }

		if !floatEquals(sd, pair.sd) {
			t.Error("Expected: ", pair.sd, "Received: ", sd)
		}
	}
}

// Test the variance function
func TestVariance(t *testing.T) {

	_, err := Variance([]float64{})

  if err == nil {
    t.Error("Expected error different from nil")
  }

	_, err = Variance([]float64{1})

  if err == nil {
    t.Error("Expected error different from nil")
  }

  // Table tests
	var tTests = []struct {
		x        []float64
		variance float64
	}{
  	{[]float64{0, 50}, 1250},
		{[]float64{10, 10, 10}, 0},
		{[]float64{10, 20, 30}, 100},
		{[]float64{0, 10, 15, 20, 35}, 167.50},
  }

	// Test with all values in the table
	for _, pair := range tTests {
		variance, err := Variance(pair.x)

    if err != nil {
			t.Error("Expected: nil Received: ", err)
    }

		if !floatEquals(variance, pair.variance) {
			t.Error("Expected: ", pair.variance, "Received: ", variance)
		}
	}
}

// Test the covariance function
func TestCovariance(t *testing.T) {

	_, err := Covariance([]float64{1, 2}, []float64{3})

  if err == nil {
    t.Error("Slices with different sizes. Error should not be nil.")
  }

	_, err = Covariance([]float64{}, []float64{})

  if err == nil {
    t.Error("Empty slices. Error should not be nil.")
  }

  // Table tests
	var tTests = []struct {
		x   []float64
    y   []float64
		res float64
	}{
    {[]float64{1, 2, 3, 4}, []float64{5, 6, 7, 8}, 1.66666667},
		{[]float64{45, 65, 98}, []float64{86, 52, 27}, -772.00000},
  }

	// Test with all values in the table
	for _, pair := range tTests {
		res, err := Covariance(pair.x, pair.y)

		if !floatEquals(res, pair.res) {
			t.Error("Expected: ", pair.res, "Received: ", res)
		}

		if err != nil {
			t.Error("Received: ", err)
		}
	}
}

// Test the covariance matrix function
func TestCovarianceMatrix(t *testing.T) {

  // Test with empty matrix
  var invalidMatrix [][]float64
  covMatrix, covError := CovarianceMatrix(invalidMatrix)

  if covError == nil {
    t.Error("Expected error different from nil")
  }

  if covMatrix != nil {
    t.Error("Expected covariance matrix equal to nil")
  }

  // Test with first row empty
  invalidMatrix = append(invalidMatrix, []float64{})
  covMatrix, covError = CovarianceMatrix(invalidMatrix)

  if covError == nil {
    t.Error("Expected error different from nil")
  }

  if covMatrix != nil {
    t.Error("Expected covariance matrix equal to nil")
  }

  // Test with different sizes of rows
  invalidMatrix = nil
  invalidMatrix = append(invalidMatrix, []float64{1, 2, 3})
  invalidMatrix = append(invalidMatrix, []float64{4, 5})
  covMatrix, covError = CovarianceMatrix(invalidMatrix)

  if covError == nil {
    t.Error("Expected error different from nil")
  }

  if covMatrix != nil {
    t.Error("Expected covariance matrix equal to nil")
  }

  var inputMatrix1 [][]float64
  inputMatrix1 = append(inputMatrix1, []float64{1, 2, 3})
  inputMatrix1 = append(inputMatrix1, []float64{6, 5, 4})
  inputMatrix1 = append(inputMatrix1, []float64{7, 8, 9})

  var outputMatrix1 [][]float64
  outputMatrix1 = append(outputMatrix1, []float64{ 1, -1,  1})
  outputMatrix1 = append(outputMatrix1, []float64{-1,  1, -1})
  outputMatrix1 = append(outputMatrix1, []float64{ 1, -1,  1})

  var inputMatrix2 [][]float64
  inputMatrix2 = append(inputMatrix2, []float64{14, 35, 97, 54, 15})
  inputMatrix2 = append(inputMatrix2, []float64{87, 24, 41, 65, 89})
  inputMatrix2 = append(inputMatrix2, []float64{12, 33, 76, 87, 43})
  inputMatrix2 = append(inputMatrix2, []float64{76, 34, 19, 11, 51})

  var outputMatrix2 [][]float64
  outputMatrix2 = append(outputMatrix2, []float64{ 1181.5, -569.5,  811.25, -689.25})
  outputMatrix2 = append(outputMatrix2, []float64{ -569.5,  811.2,  -231.8,  442.95})
  outputMatrix2 = append(outputMatrix2, []float64{ 811.25, -231.8,   956.7, -740.05})
  outputMatrix2 = append(outputMatrix2, []float64{-689.25, 442.95, -740.05,   679.7})

  var inputMatrix3 [][]float64
  inputMatrix3 = append(inputMatrix3, []float64{1, 2, 3})
  inputMatrix3 = append(inputMatrix3, []float64{4, 5, 6})
  inputMatrix3 = append(inputMatrix3, []float64{7, 8, 9})

  var outputMatrix3 [][]float64
  outputMatrix3 = append(outputMatrix3, []float64{1, 1, 1})
  outputMatrix3 = append(outputMatrix3, []float64{1, 1, 1})
  outputMatrix3 = append(outputMatrix3, []float64{1, 1, 1})

  // Table tests
  var tTests = []struct {
    inputMatrix  [][]float64
    outputMatrix [][]float64
  }{
    {inputMatrix1, outputMatrix1},
    {inputMatrix2, outputMatrix2},
    {inputMatrix3, outputMatrix3},
  }

  // Test with all values in the table
  for _, pair := range tTests {
    covMatrix, covError := CovarianceMatrix(pair.inputMatrix)

    if covError != nil {
      t.Error("Received: ", covError)
    }

    if !matrixEquals(covMatrix, pair.outputMatrix) {
      t.Error("Expected: ", pair.outputMatrix, "Received: ", covMatrix)
    }
  }
}
