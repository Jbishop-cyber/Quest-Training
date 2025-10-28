package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0 // negative numbers have no real square root
	}

	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return i // found the integer square root
		}
	}
	return 0 // no integer square root
}
