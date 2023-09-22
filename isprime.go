package piscine

func IsPrime(num int) bool {
	if num <= 1 {
		return false // Numbers less than or equal to 1 are not prime
	}

	if num <= 3 {
		return true // 2 and 3 are prime
	}

	if num%2 == 0 || num%3 == 0 {
		return false // Multiples of 2 and 3 are not prime
	}

	i := 5
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false // If divisible by i or i+2, not prime
		}
		i += 6
	}

	return true // It's prime
}
