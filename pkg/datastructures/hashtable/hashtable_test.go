package hashtable_test

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	test := "B"

	for i, a := range test {
		fmt.Println(i, a, rune(a-'A'+1))
	}
}
