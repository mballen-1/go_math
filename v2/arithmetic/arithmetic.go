package arithmetic

import "math"

func Factorial(n int) int {
	var f int = 1
	for i := 2; i <= n; i++ {
		f *= i
	}
	return f
}

// IsPrime Checks if a number is prime or not
func IsPrime(num int) bool {
	for i := 2; i < int(num/2); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// CylinderVolume
func CylinderVolume(r, h float64) float64 {
	return math.Pi * math.Pow(r, 2.) * h
}
