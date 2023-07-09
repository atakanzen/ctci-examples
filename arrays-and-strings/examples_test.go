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

func BenchmarkOnePointOne(b *testing.B) {

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
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arraysandstrings.HasAllUniqueChars(tt.input)
			}
		})
	}
}

func TestExampleOnePointTwoWithHashmap(t *testing.T) {
	tests := []struct {
		inputOne, inputTwo string
		want               bool
	}{
		{"qwert", "trewq", true},
		{"abba", "baba", true},
		{"dafgq", "fqgad", true},
		{"dog", "god    ", false},
		{"a", "ba", false},
		{"abc", "abb", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s and %s", tt.inputOne, tt.inputTwo)
		t.Run(testName, func(t *testing.T) {
			actual := arraysandstrings.CheckIfPermutation(tt.inputOne, tt.inputTwo)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestExampleOnePointTwoWithRunes(t *testing.T) {
	tests := []struct {
		inputOne, inputTwo string
		want               bool
	}{
		{"qwert", "trewq", true},
		{"abba", "baba", true},
		{"dafgq", "fqgad", true},
		{"dog", "god    ", false},
		{"a", "ba", false},
		{"abc", "abb", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s and %s", tt.inputOne, tt.inputTwo)
		t.Run(testName, func(t *testing.T) {
			actual := arraysandstrings.CheckIfPermutationWithRunes(tt.inputOne, tt.inputTwo)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func FuzzOnePointTwo(f *testing.F) {
	f.Fuzz(func(t *testing.T, inputOne, inputTwo string) {
		arraysandstrings.CheckIfPermutation(inputOne, inputTwo)
	})
}

func BenchmarkOnePointTwoWithHashmap(b *testing.B) {

	tests := []struct {
		inputOne, inputTwo string
		want               bool
	}{
		{"qwert", "trewq", true},
		{"abba", "baba", true},
		{"dafgq", "fqgad", true},
		{"a", "ba", false},
		{"abc", "abb", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s and %s", tt.inputOne, tt.inputTwo)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arraysandstrings.CheckIfPermutation(tt.inputOne, tt.inputTwo)
			}
		})
	}
}

func BenchmarkOnePointTwoWithRunes(b *testing.B) {

	tests := []struct {
		inputOne, inputTwo string
		want               bool
	}{
		{"qwert", "trewq", true},
		{"abba", "baba", true},
		{"dafgq", "fqgad", true},
		{"a", "ba", false},
		{"abc", "abb", false},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s and %s", tt.inputOne, tt.inputTwo)
		b.Run(testName, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arraysandstrings.CheckIfPermutationWithRunes(tt.inputOne, tt.inputTwo)
			}
		})
	}
}
