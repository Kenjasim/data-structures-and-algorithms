// package insertionsort

// // Sort - Uses selection sort algorithm to sort a list and return the sorted list
// func Sort(values []int) []int {
// 	startIndex := 1

// 	for {

// 		startItem := values[startIndex]

// 		for index, item := range values {

// 			if index > startIndex {
// 				continue
// 			}

// 			if item < startItem {
// 				values[index+1] = item

// 			}

// 		}

// 		startIndex++
// 	}

// 	return values
// }
