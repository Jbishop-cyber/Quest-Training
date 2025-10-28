package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false // 1 and negative numbers are not prime
	}
	if nb == 2 {
		return true // 2 is prime
	}
	if nb%2 == 0 {
		return false // even numbers >2 are not prime
	}

	for i := 3; i*i <= nb; i += 2 { // check only odd divisors
		if nb%i == 0 {
			return false // divisible by i, not prime
		}
	}
	return true
}
