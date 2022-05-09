package linearsearch

import (
	"github.com/Kenjasim/data-structures-and-algorithms/pkg/algorithms"
)

// Search - Runs a linear search algorithm on a list of values, returns the index
func Search(values []int, searchValue int) (index int, err error) {
	for index, item := range values {
		if item == searchValue {
			return index, nil
		}
	}

	return 0, &algorithms.ItemNotInArray{Item: searchValue}
}
