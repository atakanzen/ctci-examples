package arraysandstrings

import (
	"math"
	"strconv"
	"strings"
	"unicode"
)

// --------------------------------------------------------------------------------------------------------------

// 1.1: Implement an algorithm to determine if a string has all unique characters, what if you cannot use additional data structures?

// We can use a hashmap to create a frequency analysis, by iterating on each letter and see if it exists in the
// map, if yes increment the counter, if not assign it to the map. Without additional DS I'd need to go with
// O(N2) since I need to check each pair one by one.

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

// What I missed from the Book?

/*
	- Asking whether string is ASCII or Unicode to the interviewer
		- This effects the storage size of the alphabet, which is 128 possible values in ASCII, and 256 in extended ASCII
	- I couldn't think about returning directly false if the string length exceeds the possible number of unique chars
	- Asking about whether we're allowed to modify the string to sort it and linearly checking neighboring chars
*/

// --------------------------------------------------------------------------------------------------------------

// 1.2 Check Permutation: Given two strings, write a method to decide if one is a permutation of the other

// abac , baac , caab
// In this case permutations are same length, so we can immediately return false if lengths are not equal
// We can make use of a hashmap to count the occurrences and then see if they match in terms of number of occurrences and the characters

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

// What I missed from the book?

/*
	- Not necessarily about the book, but I could have used an integer array to map each rune within a predefined integer list, i.e. alphabet. See code below
*/

// Improvement

func CheckIfPermutationWithRunes(stringOne, stringTwo string) bool {
	// Here making ASCII array, depends on the incoming string so derive from interviewer
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

// Input: "My John Smith    ", 13
// Output: "Mr%20John%20Smith"

// Seeing this example, I would ask if it's possible to trim the excessive empty spaces around the input
// We can then just loop in the string, check if it's white space, then replace it with %20

// Hint One: Easiest to modify strings if going from end to start

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

// To be a permutation of a palindrome a string can have no more than one character that is odd!

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

// pale, ple -> true
// pales, pale -> true
// pale, bale -> true
// pale, bake -> false

// Zero edit means that the two strings are identical (BASE CASE)
// Number of characters if the difference is 1 that covers Insert/Remove
// Replacement -> lengths should be same, and we can loop over the first String compare the characters on iteration index

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

// Maybe a string builder usage? Not sure if in interviews these are ok. IT'S OK TO GET BETTER SPACE COMPLEXITY
// We can also count the final compressed string length and check before creating the actual string to not do unnecessary work. This will bring more code, almost duplicate, however it may be efficient. I'll benchmark these two

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
