package leetcode

// assuming both strings are same utf-8 encoded, and no
// UNICODE special characters are present
func ValidAnagram(s1, s2 string) bool {
	// if both strings are not of the same, then anagram is not possible
	if len(s1) != len(s2) {
		return false
	}

	hashMap := make(map[rune]int)
	// create the hashmap on s1
	for _, ch := range s1 {
		// if the character doesn't exist in the hashmap
		// make count = 0
		// else, increment count
		if hashMap[ch] != 0 {
			hashMap[ch] += 1
		} else {
			hashMap[ch] = 1
		}
	}
	// check 2nd string against the hashmap of 1st string
	for _, ch := range s2 {
		if hashMap[ch] == 0 {
			return false
		}
		hashMap[ch] -= 1
	}
	// check that the hashmap is perfectly consumed
	for _, ch := range s1 {
		if hashMap[ch] != 0 {
			return false
		}
	}
	return true
}
