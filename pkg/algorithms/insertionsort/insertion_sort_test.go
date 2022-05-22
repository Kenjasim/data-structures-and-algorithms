package insertionsort_test

import (
	"reflect"
	"testing"

	"github.com/Kenjasim/data-structures-and-algorithms/pkg/algorithms/insertionsort"
)

// testSort - tests the sort function
func TestSort(t *testing.T) {

	testcases := []struct {
		values         []int
		expectedValues []int
	}{
		{
			values:         []int{1, 2, 3},
			expectedValues: []int{1, 2, 3},
		},
		{
			values:         []int{2, 1, 3},
			expectedValues: []int{1, 2, 3},
		},
		{
			values:         []int{2, 1, 3, 6, 8, 10, 50, 32, 6},
			expectedValues: []int{1, 2, 3, 6, 6, 8, 10, 32, 50},
		},
	}

	for _, testcase := range testcases {
		sortedValues := insertionsort.Sort(testcase.values)

		if !reflect.DeepEqual(sortedValues, testcase.expectedValues) {
			t.Fatalf("List %v does not match expected list %v", sortedValues, testcase.expectedValues)
		}

	}

}
