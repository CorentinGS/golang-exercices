package main

func main() {
	println(lengthOfLongestSubstring("abcabcbb"))
	println(lengthOfLongestSubstring("bbbbb"))
	println(lengthOfLongestSubstring("pwwkew"))
	println(lengthOfLongestSubstring(" "))
	println(lengthOfLongestSubstring(""))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int) // hash map to store the characters in the substring
	max, left := 0, 0       // max is the max length of substring, left is the left index of the substring

	for i, c := range s {
		// if the character is already in the substring, update the left index
		if _, ok := m[c]; ok && m[c] >= left {
			// update max if the length of the substring is larger than the current max
			if max < i-left {
				max = i - left
			}
			// if the character is already in the substring, then we need to shrink the substring
			// by moving the left index to the right of the character
			left = m[c] + 1
		}
		m[c] = i // record the index of the character in the substring
	}

	// if the length of the substring is larger than the current max, update the max
	if max < len(s)-left {
		max = len(s) - left
	}

	return max
}
