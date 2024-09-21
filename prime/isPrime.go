package prime

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
