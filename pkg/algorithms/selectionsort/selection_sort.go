package selectionsort

// Sort - Uses selection sort algorithm to sort a list and return the sorted list
func Sort(values []int) []int {
	startIndex := 0
	lowestIndex := 0

	for {

		// If we have run too much then break the loop
		if startIndex >= len(values) {
			break
		}

		// Store the lowest value and its index, assume first is lowest for now
		lowestIndex = startIndex

		// See if there are any values lower than the one at the start
		for index, value := range values {
			// If the index is lower than the lowest index, we assume that this has already been seen by the list
			if index < startIndex {
				continue
			}

			if value < values[lowestIndex] {
				lowestIndex = index
			}
		}

		// If there is a value lower than the one at the first index then swap them
		if lowestIndex != startIndex {
			// Swap the values
			tempItem := values[startIndex]
			values[startIndex] = values[lowestIndex]
			values[lowestIndex] = tempItem
		}

		startIndex++
	}

	return values
}
