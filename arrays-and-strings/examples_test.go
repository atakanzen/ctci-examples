package arraysandstrings_test

import (
	"fmt"
	"testing"

	arraysandstrings "github.com/atakanzen/ctci-examples/arrays-and-strings"
	"github.com/stretchr/testify/assert"
)

func TestExampleOnePointOne(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"asdfghjkl", true},
		{"a", true},
		{"qwertyuiopq", false},
		{"qawsedrftgyhujikolpzxcvbndm", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.input)
		t.Run(testName, func(t *testing.T) {
			actual := arraysandstrings.HasAllUniqueChars(tt.input)
			assert.Equal(t, tt.want, actual)
		})
	}

}
