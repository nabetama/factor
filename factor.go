package factor

import (
	"math"
)

func factor(n int) []int {

	var primeFactors []int

	if n < 2 {
		return []int{n}
	}

	for i := 2; i < int(math.Pow(float64(n), 0.5))+1; i++ {
		for n%i == 0 {
			primeFactors = append(primeFactors, i)
			n /= i
		}
	}

	if n > 1 {
		primeFactors = append(primeFactors, n)
	}

	return primeFactors
}
