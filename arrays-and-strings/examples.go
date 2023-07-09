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
