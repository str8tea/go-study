package internal

func IsPrime(n int) bool {
	if n == 1 || n == 2 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return true
		}
	}

	return false
}
