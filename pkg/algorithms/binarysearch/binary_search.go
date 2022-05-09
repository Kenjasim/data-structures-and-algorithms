package binarysearch

import (
	"github.com/Kenjasim/data-structures-and-algorithms/pkg/algorithms"
)

// Search - Runs a binary search algorithm on a list of values,
// it expects the list to be sorted, returns the index if exists othewise returns error
func Search(values []int, searchValue int) (index int, err error) {

	searchIndex := 0

	for len(values) > 0 {
		midpoint := len(values) / 2
		midItem := values[midpoint]

		if midItem == searchValue {
			return searchIndex + midpoint, nil
		} else if len(values) == 1 {
			break
		} else if midItem < searchValue {
			values = values[midpoint:]
			searchIndex = searchIndex + midpoint
		} else {
			values = values[0:midpoint]
		}
	}

	return 0, &algorithms.ItemNotInArray{Item: searchValue}
}
