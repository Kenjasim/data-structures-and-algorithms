package hashtable

import (
	"reflect"
	"testing"
)

// TestCreateHashTable - Test the hash table factory
func TestCreateHashTable(t *testing.T) {

	ht := New()

	if reflect.TypeOf(ht.values[1]) != reflect.TypeOf(&Bucket{}) {
		t.Fatal()
	}

}

// TestBucketSearch - Test searching a bucket
func TestBucketSearch(t *testing.T) {

	testcases := []struct {
		bucket        Bucket
		keyToSearch   string
		expectedFound bool
		expectedIndex int
	}{
		{
			bucket:        Bucket{values: []*Node{{Key: "foo", Value: "bar"}}},
			keyToSearch:   "foo",
			expectedFound: true,
			expectedIndex: 0,
		},
		{
			bucket:        Bucket{values: []*Node{{Key: "bar", Value: "foo"}}},
			keyToSearch:   "foo",
			expectedFound: false,
			expectedIndex: 0,
		},
		{
			bucket:        Bucket{values: []*Node{{Key: "bar", Value: "foo"}, {Key: "foo", Value: "bar"}}},
			keyToSearch:   "foo",
			expectedFound: true,
			expectedIndex: 1,
		},
	}

	for _, testcase := range testcases {
		found, index := testcase.bucket.Search(testcase.keyToSearch)

		if found != testcase.expectedFound {
			t.Fatalf("Found value %T does not match expected found value %T", found, testcase.expectedFound)
		}

		if index != testcase.expectedIndex {
			t.Fatalf("Index %d does not match expected index %d", index, testcase.expectedIndex)
		}
	}

}

// TestBucketAppend - Test appending to a bucket
func TestBucketAppend(t *testing.T) {

	testcases := []struct {
		bucket         Bucket
		nodeToAppend   *Node
		expectedBucket Bucket
	}{
		{
			bucket:         Bucket{values: []*Node{{Key: "foo", Value: "bar"}}},
			nodeToAppend:   &Node{Key: "foo", Value: "var"},
			expectedBucket: Bucket{values: []*Node{{Key: "foo", Value: "var"}}},
		},
		{
			bucket:         Bucket{values: []*Node{{Key: "bar", Value: "foo"}}},
			nodeToAppend:   &Node{Key: "foo", Value: "var"},
			expectedBucket: Bucket{values: []*Node{{Key: "bar", Value: "foo"}, {Key: "foo", Value: "var"}}},
		},
	}

	for _, testcase := range testcases {
		testcase.bucket.Append(testcase.nodeToAppend)

		for index, item := range testcase.bucket.values {
			expectedItem := testcase.expectedBucket.values[index]
			if item.Key != expectedItem.Key || item.Value != expectedItem.Value {
				t.Fatalf("Value in bucket %v, %v not deeply equal to value %v in bucket %v", testcase.bucket.values, item, expectedItem, testcase.expectedBucket.values)
			}
		}
	}

}

// TestHashTablePut - Test putting to a hash table
func TestHashTablePut(t *testing.T) {

	testcases := []struct {
		nodeToAppend   *Node
		expectedBucket Bucket
		expectedIndex  int
	}{
		{
			nodeToAppend:   &Node{Key: "foo", Value: "var"},
			expectedBucket: Bucket{values: []*Node{{Key: "foo", Value: "var"}}},
			expectedIndex:  24,
		},
	}

	for _, testcase := range testcases {

		ht := New()

		ht.Put(testcase.nodeToAppend.Key, testcase.nodeToAppend.Value)

		// If the values in the expected index arent stored propperly then return nil
		if ht.values[testcase.expectedIndex] == nil {
			t.Fatal()
		}

		// Check if the item is in the hashtable
		for index, item := range ht.values[testcase.expectedIndex].values {
			expectedItem := testcase.expectedBucket.values[index]
			if item.Key != expectedItem.Key || item.Value != expectedItem.Value {
				t.Fatalf("Value in bucket %v, %v not deeply equal to value %v in bucket %v", ht.values[testcase.expectedIndex].values, item, expectedItem, testcase.expectedBucket.values)
			}
		}
	}

}
