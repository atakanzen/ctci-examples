package arraysandstrings

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

// 1.1: Implement an algorithm to determine if a string has all unique characters, what if you cannot use additional data structures?

func HasAllUniqueChars(input string) bool {
	set := make(map[rune]bool, 0)

	for _, ch := range input {
		if _, ok := set[ch]; ok {
			return false
		}
		set[ch] = true
	}

	return true
}

// 1.2 Check Permutation: Given two strings, write a method to decide if one is a permutation of the other

// Time complexity is O(A + B * N), dropping the constants yields O(N)

func CheckIfPermutation(stringOne, stringTwo string) bool {
	if len(stringOne) != len(stringTwo) {
		return false
	}

	// O(A + B) where A is stringOne's length and B is stringTwo's length
	frequencyMapForOne := getFrequencyMap(stringOne)
	frequencyMapForTwo := getFrequencyMap(stringTwo)

	// O(N)
	for keyOne, valOne := range frequencyMapForOne {
		if valTwo, ok := frequencyMapForTwo[keyOne]; ok {
			if valOne != valTwo {
				return false
			}
		}
	}

	return true
}

func getFrequencyMap(input string) map[rune]int {
	frequencyMap := make(map[rune]int, 0)

	for _, ch := range input {
		if val, ok := frequencyMap[ch]; ok {
			frequencyMap[ch] = val + 1
			continue
		}

		frequencyMap[ch] = 1
	}

	return frequencyMap
}

// Improved

func CheckIfPermutationWithRunes(stringOne, stringTwo string) bool {
	// Here making ASCII array, depends on the incoming string, ask about it
	letters := make([]int, 128)
	for _, ru := range stringOne {
		// We count each rune occurence in stringOne
		// Runes are integers, and can be used as indices
		letters[ru]++
	}

	for _, ru := range stringTwo {
		// We then decrease each rune occurrence by one, if it reaches -1 that means the amount of characters are not equal between two string, meaning they are not permutations
		letters[ru]--
		if letters[ru] < 0 {
			return false
		}
	}

	return true
}

// --------------------------------------------------------------------------------------------------------------

// 1.3 URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string has sufficient space at the end to hold the additional characters, and that you are give the true length of the string. Perform operation in place if possible

func URLify(input string, length int) string {
	space := 0
	r := []rune{}
	for i := 0; i < length; i++ {
		if string(input[i]) == " " {
			space++
		}
	}
	for _, v := range input {
		r = append(r, v)
	}
	index := space*2 + length
	for i := length - 1; i >= 0; i-- {
		if r[i] == ' ' {
			r[index-1] = '0'
			r[index-2] = '2'
			r[index-3] = '%'
			index = index - 3
		} else {
			r[index-1] = r[i]
			index--
		}
	}

	return string(r)
}

// --------------------------------------------------------------------------------------------------------------

// 1.4 Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome. Palindrome doesn't need to be just limited to dictionary words.

func CheckIfPalindromePermutation(input string) bool {
	letters := make([]int, 128)
	for _, ru := range input {
		if string(ru) != " " {
			letters[unicode.ToLower(ru)]++
		}
	}

	oddFound := false

	for _, letterCount := range letters {
		if letterCount%2 == 1 {
			if oddFound {
				return false
			}
			oddFound = true
		}
	}

	return true
}

// --------------------------------------------------------------------------------------------------------------

// 1.5 One Away: There are three types of edits that can be performed on strings:
// 	- Insert a character
// 	- Remove a character
// 	- Replace a character
//    Given two strings, write a function to check if they are one edit (or zero edits) away

func IsOneAwayEdited(inputOne, inputTwo string) bool {
	// Zero edit
	if inputOne == inputTwo {
		return true
	}

	// Insert/Remove
	if math.Abs(float64(len(inputOne)-len(inputTwo))) == 1 {
		return true
	}

	// Replacement
	changeFound := false
	for i, ru := range inputOne {
		if rune(inputTwo[i]) != ru {
			if changeFound {
				return false
			}
			changeFound = true
		}
	}

	return true
}

// 1.6 String Compression: Implement a function to perform basic string compression using the counts of repeated characters. E.g. string aabcccccaaa would become a2b1c5a3. If the compressed string would not become smaller than the original string. You can assume the string has only uppercase and lowercase letters (a-z)

// A BIT FASTER IN TIME COMPLEXITY BUT WORSE IN SPACE COMPLEXITY
func StringCompression(originalString string) string {
	counter := 0
	sb := new(strings.Builder)

	for i, ru := range originalString {
		counter++
		if i+1 >= len(originalString) || ru != rune(originalString[i+1]) {
			sb.WriteString(string(ru))
			sb.WriteString(strconv.Itoa(counter))
			counter = 0
		}
	}

	compressed := sb.String()

	if len(compressed) < len(originalString) {
		return compressed
	} else {
		return originalString
	}
}

// FASTER TIME COMPLEXITY IF NO NEED TO COMPRESS, MUCH BETTER SPACE COMPLEXITY
func StringCompressionPreliminaryCheck(originalString string) string {
	finalCompressedLength := countCompressionLength(originalString)
	// Preliminary check and return
	if finalCompressedLength > len(originalString) {
		return originalString
	}

	counter := 0
	sb := new(strings.Builder)
	sb.Grow(finalCompressedLength)

	for i, ru := range originalString {
		counter++
		if i+1 >= len(originalString) || ru != rune(originalString[i+1]) {
			sb.WriteString(string(ru))
			sb.WriteString(strconv.Itoa(counter))
			counter = 0
		}
	}

	compressed := sb.String()

	if len(compressed) < len(originalString) {
		return compressed
	} else {
		return originalString
	}
}

func countCompressionLength(input string) int {
	compressedLength := 0
	counter := 0
	for i, ru := range input {
		counter++
		if i+1 >= len(input) || ru != rune(input[i+1]) {
			compressedLength += 1 + len(strconv.Itoa(counter))
			counter = 0
		}
	}
	return compressedLength
}

// 1.7 Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes (32bits), write a method to rotate the image by 90 degrees, in place?

func RotateImageBy90Degrees(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix) != len(matrix[0]) {
		return matrix
	}

	n := len(matrix)
	// Amount of layers is equal to the half of matrix size
	for layer := 0; layer < n/2; layer++ {
		// Boundaries of the layer
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			// Saving top first
			top := matrix[first][i]
			// Left -> Top
			matrix[first][i] = matrix[last-offset][first]
			// Bottom -> Left
			matrix[last-offset][first] = matrix[last][last-offset]
			// Right -> Bottom
			matrix[last][last-offset] = matrix[i][last]
			// Top -> Right
			matrix[i][last] = top
		}
	}

	return matrix
}

// 1.8 Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0

func ZeroMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	zeroFound := false
	rowPos := 0
	colPos := 0

	// O(M*N)
	for rowIndex, row := range matrix {
		if zeroFound {
			break
		}

		for colIndex, _ := range row {
			if matrix[rowIndex][colIndex] == 0 {
				zeroFound = true
				rowPos = rowIndex
				colPos = colIndex
				break
			}
		}
	}

	if zeroFound {
		for i := 0; i < len(matrix[rowPos]); i++ {
			matrix[rowPos][i] = 0
		}

		for i := 0; i < len(matrix); i++ {
			matrix[i][colPos] = 0
		}
	}

	return matrix
}

// 1.9 String Rotation: Assume you have a method isSubstring which checks if one word is a substring of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one call to isSubstring

func IsStringRotation(s1, s2 string) bool {
	if len(s1) == len(s2) && len(s1) != 0 {
		s1s1 := s1 + s1
		return strings.Contains(s1s1, s2)
	}

	return false
}
