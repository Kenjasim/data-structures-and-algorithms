package linearsearch_test

import (
	"errors"
	"testing"

	"github.com/Kenjasim/data-structures-and-algorithms/pkg/algorithms"
	"github.com/Kenjasim/data-structures-and-algorithms/pkg/algorithms/linearsearch"
)

// testSearch - tests the search function
func TestSearch(t *testing.T) {

	testcases := []struct {
		values        []int
		searchValue   int
		expectedIndex int
		expectedErr   error
	}{
		{
			values:        []int{1, 2, 3},
			searchValue:   2,
			expectedIndex: 1,
			expectedErr:   nil,
		},
		{
			values:        []int{2, 5, 7},
			searchValue:   7,
			expectedIndex: 2,
			expectedErr:   nil,
		},
		{
			values:        []int{2, 5, 7},
			searchValue:   6,
			expectedIndex: 0,
			expectedErr:   &algorithms.ItemNotInArray{Item: 6},
		},
	}

	for _, testcase := range testcases {
		index, err := linearsearch.Search(testcase.values, testcase.searchValue)

		if err != nil && !errors.As(err, &testcase.expectedErr) {
			t.Fatalf("Error %v does not match expected error %v", err, testcase.expectedErr)
		}

		if index != testcase.expectedIndex {
			t.Fatalf("Index %d does not match expected index %d", index, testcase.expectedIndex)
		}
	}

}
