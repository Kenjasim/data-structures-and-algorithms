package hashtable

import (
	"fmt"

	"github.com/Kenjasim/data-structures-and-algorithms/pkg/utils/hash"
)

type KeyNotFound struct {
	Key string
}

func (e *KeyNotFound) Error() string {
	return fmt.Sprintf("key %s not found in hashtable", e.Key)
}

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
		values[index] = &Bucket{values: []*Node{}}
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
		if node != nil && node.Key == key {
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

// Get - gets the item into the hash table
func (h *HashTable) Get(key string) (value string, err error) {

	// Get the hash table value of the key to get the index
	hashIndex := hash.HashString(key)

	// Access the correct index and append the values
	bucket := h.values[hashIndex]

	// Search for the item in the bucket
	found, index := bucket.Search(key)

	// Return error is the item is not found
	if !found {
		return "", &KeyNotFound{Key: key}
	}

	// Return the value found
	return bucket.values[index].Value, nil

}

// Delete the item from the hashtable
func (h *HashTable) Delete(key string) (err error) {

	// Get the hash table value of the key to get the index
	hashIndex := hash.HashString(key)

	// Access the correct index and append the values
	bucket := h.values[hashIndex]

	// Search for the item in the bucket
	found, index := bucket.Search(key)

	// Return error is the item is not found
	if !found {
		return &KeyNotFound{Key: key}
	}

	// Delete the value
	bucket.values = append(bucket.values[:index], bucket.values[index+1:]...)
	return nil
}
