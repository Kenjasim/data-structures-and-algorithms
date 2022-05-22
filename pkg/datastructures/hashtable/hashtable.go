package hashtable

// HashTable - Contains the list of items which will be stored
type HashTable struct {
	values []string
}

type Item struct {
	items []string
}

// New - creates and returns an empty hashtable
func New() *HashTable {
	return &HashTable{values: make([]string, 100)}
}

// // Put - puts the item into the hash table
// func (h *HashTable) Put(item string) error {

// }
