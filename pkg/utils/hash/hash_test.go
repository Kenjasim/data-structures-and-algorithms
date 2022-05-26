package hash_test

import (
	"testing"

	"github.com/Kenjasim/data-structures-and-algorithms/pkg/utils/hash"
)

// testSearch - tests the search function
func TestHashString(t *testing.T) {

	testcases := []struct {
		hashString   string
		expectedHash int
	}{
		{
			hashString:   "foo",
			expectedHash: 24,
		},
		{
			hashString:   "bar",
			expectedHash: 9,
		},
		{
			hashString:   "foobar",
			expectedHash: 33,
		},
		{
			hashString:   "foo bar",
			expectedHash: 65,
		},
		{
			hashString:   "a",
			expectedHash: 97,
		},
		{
			hashString:   "A",
			expectedHash: 65,
		},
		{
			hashString:   "!",
			expectedHash: 33,
		},
		{
			hashString:   "H! h0w_Ar3_Y0u?",
			expectedHash: 45,
		},
	}

	for _, testcase := range testcases {
		hashValue := hash.HashString(testcase.hashString)

		if hashValue != testcase.expectedHash {
			t.Fatalf("Hash %d does not match expected hash %d", hashValue, testcase.expectedHash)
		}
	}

}
