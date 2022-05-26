package hash

// HashString - Hash function to return a value between 0 and 99 given a string
func HashString(s string) int {
	length := len(s)
	sumation := 0

	// Sum the values of the string
	for i := 0; i < length; i++ {
		sumation += int(s[i])
	}

	hash := sumation % 100
	return hash
}
