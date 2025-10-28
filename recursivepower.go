package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0 // negative powers return 0
	}
	if power == 0 {
		return 1 // base case: nb^0 = 1
	}
	return nb * RecursivePower(nb, power-1) // recursive step
}
