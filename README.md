# Statistics

[![Build Status](https://travis-ci.org/kelvins/statistics.svg?branch=master)](https://travis-ci.org/kelvins/statistics)
[![Build Status](https://circleci.com/gh/kelvins/statistics.svg?style=shield&circle-token=:circle-token)](https://circleci.com/gh/kelvins/statistics)
[![Coverage Status](https://coveralls.io/repos/github/kelvins/statistics/badge.svg?branch=master)](https://coveralls.io/github/kelvins/statistics?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kelvins/statistics)](https://goreportcard.com/report/github.com/kelvins/statistics)
[![GoDoc](https://godoc.org/github.com/kelvins/statistics?status.svg)](https://godoc.org/github.com/kelvins/statistics)
[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](LICENSE)

Go package that provides some statistical functions.

## Installation

Use:

    $ go get github.com/kelvins/statistics

## Methods

- **Mean**: The mean formula basically says: “Add up all the numbers and then divide by how many there are”.

<center>![Mean](http://i.imgur.com/nepiG0s.png)</center>

- **Standard Deviation**: The Standard Deviation (SD) of a data set is a measure of how spread out the data is. Basically, the standard deviation is defined as “the average distance from the mean of the data set to a point”. As Smith explains, the way to calculate it is to compute the squares of the distance from each data point to the mean of the set, add them all up,
divide by `(n - 1)`, and take the positive square root.

<center>![Mean](http://i.imgur.com/A6nUpGU.png)</center>

- **Note**: The `(n - 1)` expression should be used when we have a sample of the data set. If we want to calculate the SD for an entire population we should divide by `n` instead of `(n - 1)`.

- **Variance**: variance is another measure of the spread of data in a data set. In fact, it is almost identical to the standard deviation. The variance is simply the standard deviation squared, in both the symbol and the formula (there is no square root in the formula for variance).

<center>![Mean](http://i.imgur.com/qZu5to3.png)</center>

- **Covariance**: standard deviation and variance only operate on 1 dimension, so that you could only calculate the standard deviation for each dimension of the data set independently of the other dimensions. However, it is useful to have a similar measure to find out how much the dimensions vary from the mean with respect to each other. Covariance is such a measure. Covariance is always measured between 2 dimensions. If you calculate the covariance between one dimension and itself, you get the variance. So, if you had a 3-dimensional data set `(x, y, z)`, then you could measure the covariance between the `x` and `y` dimensions, the `x` and `z` dimensions, and the `y` and `z` dimensions. So basically the covariance formula says: “for each data item, multiply the difference between the `x` value and the mean of `x`, by the the difference between the `y` value and the mean of `y`. Add all these up, and divide by `(n - 1)`”.

<center>![Mean](http://i.imgur.com/Wt1c76u.png)</center>

- **Covariance Matrix**: Recall that covariance is always measured between 2 dimensions. If we have a data set with more than 2 dimensions, there is more than one covariance measurement that can be calculated. A useful way to get all the possible covariance values between all the different dimensions is to calculate them all and put them in a matrix. So, supposing that we have a 3-dimensional data set `(x, y, z)` the covariance matrix will have 3 rows and 3 columns, and will look like this:

<center>![Mean](http://i.imgur.com/H1xM95q.png)</center>

**Notes**: down the main diagonal, we can see that the covariance value is between one of the dimensions and itself. It means that these are the variances for that dimension. The other point is that since `cov(x, y) = cov(y, x)`, the matrix is symmetrical about the main diagonal.

## Documentation


You can access the [documentation][2] by the [GoDoc][3] website.

## License

This project was created under the [MIT license][1], feel free to contribute in any way.

## Contributing

1. Create an issue (optional)
2. Fork the repo
3. Make your changes
4. Commit your changes (`git commit -am 'Some cool feature'`)
5. Push to the branch (`git push origin master`)
6. Create a new Pull Request

## References

**Lindsay I Smith**. A tutorial on Principal Components Analysis. 2002. Link: [http://www.cs.otago.ac.nz/cosc453/student_tutorials/principal_components.pdf][4]

  [1]: LICENSE
  [2]: https://godoc.org/github.com/kelvins/statistics
  [3]: https://godoc.org/
  [4]: http://www.cs.otago.ac.nz/cosc453/student_tutorials/principal_components.pdf
