Statistics
==========================

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/kelvins/statistics)](https://goreportcard.com/report/github.com/kelvins/statistics)
[![GoDoc](https://godoc.org/github.com/kelvins/statistics?status.svg)](https://godoc.org/github.com/kelvins/statistics)

Go package that provides some statistical functions.

- Mean
 - ![mean](http://www.sciweavers.org/tex2img.php?eq=%5Cbar%7BX%7D%20%3D%20%5Cfrac%7B%5Csum_%7Bi%3D1%7D%5E%7Bn%7D%20X_i%7D%7Bn%7D&bc=White&fc=Black&im=png&fs=12&ff=arev&edit=0)
- Standard Deviation
 - ![mean](http://www.sciweavers.org/tex2img.php?eq=s%20%3D%20%5Csqrt%7B%20%5Cfrac%7B%5Csum_%7Bi%3D1%7D%5E%7Bn%7D%20%28X_i%20-%20%5Cbar%7BX%7D%29%5E2%7D%7B%28n-1%29%7D%20%7D&bc=White&fc=Black&im=png&fs=12&ff=arev&edit=0)
- Variance
 - ![mean](http://www.sciweavers.org/tex2img.php?eq=s%5E2%20%3D%20%5Cfrac%7B%5Csum_%7Bi%3D1%7D%5E%7Bn%7D%20%28X_i%20-%20%5Cbar%7BX%7D%29%5E2%7D%7B%28n-1%29%7D&bc=White&fc=Black&im=png&fs=12&ff=arev&edit=0)
- Covariance
 - ![mean](http://www.sciweavers.org/tex2img.php?eq=cov%28X%2C%20Y%29%20%3D%20%5Cfrac%7B%5Csum_%7Bi%3D1%7D%5E%7Bn%7D%20%28X_i%20-%20%5Cbar%7BX%7D%29%20%28Y_i%20-%20%5Cbar%7BY%7D%29%7D%7B%28n-1%29%7D&bc=White&fc=Black&im=png&fs=12&ff=arev&edit=0)
- Covariance Matrix
 - Example Covariance Matrix 3x3: ![mean](http://www.sciweavers.org/tex2img.php?eq=C%20%3D%20%5Cbegin%7Bpmatrix%7D%0A%20cov%28x%2C%20x%29%20%26%20cov%28x%2C%20y%29%20%26%20cov%28x%2C%20z%29%20%5C%5C%20%0A%20cov%28y%2C%20x%29%20%26%20cov%28y%2C%20y%29%20%26%20cov%28y%2C%20z%29%20%5C%5C%20%0A%20cov%28z%2C%20x%29%20%26%20cov%28z%2C%20y%29%20%26%20cov%28z%2C%20z%29%0A%5Cend%7Bpmatrix%7D&bc=White&fc=Black&im=png&fs=12&ff=arev&edit=0)
