package piscine

// RecursiveFactorial calculates the factorial of an integer recursively.
// It returns 0 for negative numbers, numbers causing overflow (>= 21), or if the result is 0 (though factorials are always >= 1).
func RecursiveFactorial(nb int) int {
	// Base case 1: Handle error conditions (negative numbers or overflow trigger).
	// Factorial is only defined for non-negative integers.
	// 20! is the largest factorial that fits in a standard 64-bit integer (int).
	// Inputs >= 21 will cause an overflow, which is treated as an error.
	if nb < 0 || nb >= 21 {
		return 0
	}

	// Base case 2: The factorial of 0 is 1.
	if nb == 0 {
		return 1
	}

	// Recursive step: n! = n * (n-1)!
	// We multiply the current number by the factorial of the number right below it.
	return nb * RecursiveFactorial(nb-1)
}
