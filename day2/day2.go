package day2

// CountContainsTwo gets the number of values that have 
// exactly two matching characters
func CountContainsTwo(values []string) int {
	count := 0
	for _, value := range values {
		if CheckContainsTwo(value) {
			count++
		}
	}
	return count
}

// CountContainsThree gets the number of values that have 
// exactly three matching characters
func CountContainsThree(values []string) int {
	count := 0
	for _, value := range values {
		if CheckContainsThree(value) {
			count++
		}
	}
	return count
}

// CheckContainsTwo checks a string for exactly two matching characters
func CheckContainsTwo(value string) bool {
	counts := make(map[rune]int)

	for _, char := range value {
		counts[char]++
	}

	for _, count := range counts {
		if count == 2 {
			return true
		}
	}

	return false
}

// CheckContainsThree checks a string for exactly two matching characters
func CheckContainsThree(value string) bool {
	counts := make(map[rune]int)

	for _, char := range value {
		counts[char]++
	}

	for _, count := range counts {
		if count == 3 {
			return true
		}
	}

	return false
}

// GetCorrectID finds the "correct" ID according to the closest matching values
func GetCorrectID(values []string) string {
	for i, valueA := range values {
		for j, valueB := range values {
			if i == j {
				continue
			}

			diffIdx := 0
			diffCount := 0
			for idx, charA := range valueA {
				if charA != rune(valueB[idx]) {
					diffIdx = idx
					diffCount++
				}
			}
			
			if diffCount == 1 {
				return valueA[:diffIdx] + valueA[diffIdx+1:]
			}
		}
	}

	return ""
}
