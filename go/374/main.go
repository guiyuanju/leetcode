package main

func main() {
}

func guessNumber(n int) int {
	left := 0
	right := n
	for left <= right {
		mid := left + (right - left) / 2
		g := guess(mid)
		if g == 0 {
			return mid
		} else if g > 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
