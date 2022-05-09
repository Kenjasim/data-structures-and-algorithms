package bubblesort

// Sort - Uses bubblesort algorithm to sort a list and return the sorted list
func Sort(values []int) []int {
	sorted := false

	// If we are passed an empty list or a list with only 1 item return it
	if len(values) < 2 {
		return values
	}

	for !sorted {
		// Assume sorted until we see otherwise
		sorted = true

		for index, value := range values {
			if index == len(values)-1 {
				break
			}

			// Check if the next item is smaller than the first item, if so they should swap places
			nextItem := values[index+1]
			if value > nextItem {
				values[index+1] = value
				values[index] = nextItem
				sorted = false
			}

		}
	}

	return values
}
