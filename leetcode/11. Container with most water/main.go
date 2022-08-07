package main

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	println(maxArea(height))
}

func maxArea(height []int) int {
	var area int
	left, right := 0, len(height)-1

	for left < right {
		area = max((right-left)*min(height[left], height[right]), area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return area
}

func max(i int, i2 int) int {
	if i >= i2 {
		return i
	}
	return i2
}

func min(i int, i2 int) int {
	if i <= i2 {
		return i
	}
	return i2
}
