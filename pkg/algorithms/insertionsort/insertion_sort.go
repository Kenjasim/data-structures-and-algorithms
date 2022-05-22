package insertionsort

// Sort - Uses selection sort algorithm to sort a list and return the sorted list
func Sort(values []int) []int {

	for startIndex := 1; startIndex < len(values); startIndex++ {

		startItem := values[startIndex]
		swapIndex := 0
		swap := false

		for i := startIndex - 1; i >= 0; i-- {

			item := values[i]

			if item > startItem {
				values[i+1] = item
				swapIndex = i
				swap = true
			}

		}

		if swap {
			values[swapIndex] = startItem
		}
	}

	return values
}
