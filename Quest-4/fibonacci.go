package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1 // negative index returns -1
	}
	if index == 0 {
		return 0 // base case: Fib(0) = 0
	}
	if index == 1 {
		return 1 // base case: Fib(1) = 1
	}
	return Fibonacci(index-1) + Fibonacci(index-2) // recursive step
}
