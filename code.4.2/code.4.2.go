package main

import (
	"fmt"
	"math"
)

func solveSquareEquation(a float64, b float64, c float64, p func(int, float64, float64)) {
	discFinder := func(a float64, b float64, c float64) float64 {
		return b * b - 4 * a * c
	}

	D := discFinder(a, b, c)

	switch {
	case D > 0:
		p(2, (-b + math.Sqrt(D)) / (2 * a), (-b - math.Sqrt(D)) / (2 * a))
	case D == 0:
		p(1, -b / (2 * a), 0)
	default:
		p(0, 0, 0)
	}
}

func main() {
	var a, b, c float64
	fmt.Scanf("%f %f %f", &a, &b, &c)

	solveSquareEquation(a, b, c, func (n int, x1 float64, x2 float64) {
		fmt.Println(n, x1, x2)
	})
}
