package arraysandstrings

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
