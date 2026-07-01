package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddOne(t *testing.T) {
	// var (
	// 	input    int = 1
	// 	expected int = 3
	// )

	// result := AddOne(input)
	// if result != expected {
	// 	t.Errorf("AddOne(%d) = %d; want %d", input, result, expected)
	// }

	assert.Equal(t, AddOne(1), 2, "AddOne(1) should be 2")
	assert.NotEqual(t, AddOne(1), 3, "AddOne(1) should not be 3")
}
