package hashtable

import "github.com/Kenjasim/data-structures-and-algorithms/pkg/utils/hash"

// HashTable - Contains the list of items which will be stored
type HashTable struct {
	values [100]*Bucket
}

type Bucket struct {
	values []*Node
}

type Node struct {
	Key   string
	Value string
}

// New - creates and returns an empty hashtable
func New() *HashTable {
	// Initialise the nodes for each item in the hash table
	var values [100]*Bucket

	for index := range values {
		values[index] = &Bucket{make([]*Node, 10)}
	}

	// Return the values
	return &HashTable{values: values}
}

// Append a value to the bucket
func (b *Bucket) Append(node *Node) {
	// Search to see if the key is already in the hash table
	found, index := b.Search(node.Key)

	if found {
		b.values[index] = node
		return
	}

	b.values = append(b.values, node)

}

// Search - search for a key to see if it has been stored in the hash table and return the index if found
func (b *Bucket) Search(key string) (found bool, index int) {

	for index, node := range b.values {
		if node.Key == key {
			return true, index
		}
	}

	return false, 0
}

// Put - puts the item into the hash table
func (h *HashTable) Put(key, value string) {

	// Get the hash table value of the key to get the index
	hashIndex := hash.HashString(key)

	// Access the correct index and append the values
	node := Node{Key: key, Value: value}
	bucket := h.values[hashIndex]
	bucket.Append(&node)

}
