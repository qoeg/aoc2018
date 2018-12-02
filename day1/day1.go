package day1

// CalculateFreq calculates a value based on a starting frequency
// and a single set of "frequency changes" to be used only once
func CalculateFreq(value int, changes []int) int {
	for _, change := range changes {
		value += change
	}

	return value
}

// CalibrateFreq calibrates a true frequency by continuously processing
// a value until a duplicate frequency is found (which is the true frequency)
func CalibrateFreq(value int, changes []int) int {
	results := make(map[int]int)

	for {
		for _, change := range changes {
			value += change
			results[value]++

			if (results[value] > 1) {
				return value
			}
		}
	}
}
