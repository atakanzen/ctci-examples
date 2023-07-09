package arraysandstrings

// 1.1: Implement an algorithm to determine if a string has all unique characters, what if you cannot use additional data structures?

// We can use a hashmap to create a frequency analysis, by iterating on each letter and see if it exists in the
// map, if yes increment the counter, if not assign it to the map. Without additional DS I'd need to go with
// O(N2) since I need to check each pair one by one.

func HasAllUniqueChars(input string) bool {
	frequencyTable := make(map[rune]int, 0)

	for _, ch := range input {
		if _, ok := frequencyTable[ch]; ok {
			return false
		}
		frequencyTable[ch] = 1
	}

	return true
}
