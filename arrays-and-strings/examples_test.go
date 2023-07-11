package arraysandstrings_test

import (
	"fmt"
	"testing"

	arraysandstrings "github.com/atakanzen/ctci-examples/arrays-and-strings"
	"github.com/stretchr/testify/assert"
)

// 1.1 --------------------------------------------------------------------------------------------------------

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

// 1.2 --------------------------------------------------------------------------------------------------------

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

// 1.3 --------------------------------------------------------------------------------------------------------

func TestURLify(t *testing.T) {
	tests := []struct {
		input  string
		length int
		want   string
	}{
		{
			"Mr John Smith      ",
			13,
			"Mr%20John%20Smith  ",
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.input)
		t.Run(testName, func(t *testing.T) {
			actual := arraysandstrings.URLify(tt.input, tt.length)
			assert.Equal(t, tt.want, actual)
		})
	}
}

// 1.4 --------------------------------------------------------------------------------------------------------

func TestCheckIfPermutationPalindrome(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  bool
	}{
		{
			desc:  "Tact Coa Should Return True",
			input: "Tact Coa",
			want:  true,
		},
		{
			desc:  "Qeur rique Should Return True",
			input: "Quer rique",
			want:  true,
		},
		{
			desc:  "UTQY YYTQYY Should Return False",
			input: "UTQY YYTQYY",
			want:  false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := arraysandstrings.CheckIfPalindromePermutation(tC.input)
			assert.Equal(t, tC.want, actual)
		})
	}
}

// 1.5 --------------------------------------------------------------------------------------------------------

func TestIsOneAwayEdited(t *testing.T) {
	testCases := []struct {
		desc               string
		inputOne, inputTwo string
		want               bool
	}{
		{
			desc:     "pale and Pale should return true",
			inputOne: "pale",
			inputTwo: "Pale",
			want:     true,
		},
		{
			desc:     "pale and pales should return true",
			inputOne: "pale",
			inputTwo: "pales",
			want:     true,
		},
		{
			desc:     "pale and pale should return true",
			inputOne: "pale",
			inputTwo: "pale",
			want:     true,
		},
		{
			desc:     "pale and bake should return false",
			inputOne: "pale",
			inputTwo: "bake",
			want:     false,
		},
		{
			desc:     "pale and baked too much should return false",
			inputOne: "pale",
			inputTwo: "baked too much",
			want:     false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := arraysandstrings.IsOneAwayEdited(tC.inputOne, tC.inputTwo)
			assert.Equal(t, tC.want, actual)
		})
	}
}
