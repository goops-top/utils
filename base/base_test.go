package base

import (
	"fmt"
	"testing"
)

func TestBaseTools(t *testing.T) {
	rowS := []string{"a", "Hello", "world", "hello", "World", "hello world", "Hello World", "hello world", "A", "a", "b"}
	dupS := deduplicate(rowS)
	fmt.Println(rowS)
	fmt.Println(dupS)
}
