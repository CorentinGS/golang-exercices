package main

func main() {
	println(convert("PAYPALISHIRING", 3))
}

func convert(s string, numRows int) string {
	// If numRows is 1, return s.
	if numRows == 1 {
		return s
	}

	// Create a slice of strings to hold the rows.
	array := make([][]byte, numRows)

	goDown := true // Whether we're going down or up.
	index := 0     // The index of the current row.

	// Loop through the string.
	for i := 0; i < len(s); i++ {
		// Append the current character to the current row.
		array[index] = append(array[index][:], s[i])
		if index == len(array)-1 {
			// If we're at the last row, we're going up.
			goDown = false
		}
		if index == 0 {
			// If we're at the first row, we're going down.
			goDown = true
		}
		// If we're going down, increment the index.
		if goDown {
			index++
		} else {
			// Otherwise, decrement the index.
			index--
		}
	}

	answer := ""

	// Join the rows together.
	for _, bytes := range array {
		answer += string(bytes[:]) // Convert the bytes to a string.
	}
	return answer
}
