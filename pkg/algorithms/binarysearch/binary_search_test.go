package binarysearch_test

import (
	"errors"
	"testing"

	"github.com/Kenjasim/data-structures-and-algorithms/pkg/algorithms"
	"github.com/Kenjasim/data-structures-and-algorithms/pkg/algorithms/binarysearch"
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
			values:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			searchValue:   14,
			expectedIndex: 13,
			expectedErr:   nil,
		},
		{
			values:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			searchValue:   2,
			expectedIndex: 1,
			expectedErr:   nil,
		},
		{
			values:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			searchValue:   1,
			expectedIndex: 0,
			expectedErr:   nil,
		},
		{
			values:        []int{1, 2, 3},
			searchValue:   6,
			expectedIndex: 0,
			expectedErr:   &algorithms.ItemNotInArray{Item: 6},
		},
	}

	for _, testcase := range testcases {
		index, err := binarysearch.Search(testcase.values, testcase.searchValue)

		if testcase.expectedErr == nil && err != nil {
			t.Fatalf("Error %v raised when no error expected", err)
		}

		if err != nil && !errors.As(err, &testcase.expectedErr) {
			t.Fatalf("Error %v does not match expected error %v", err, testcase.expectedErr)
		}

		if index != testcase.expectedIndex {
			t.Fatalf("Index %d does not match expected index %d", index, testcase.expectedIndex)
		}
	}

}
