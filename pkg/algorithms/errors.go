package algorithms

import "fmt"

type ItemNotInArray struct {
	Item int
}

func (e *ItemNotInArray) Error() string {
	return fmt.Sprintf("item %d not found in array", e.Item)
}
