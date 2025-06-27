package main

func main() {
	println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	println(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	var res int
	i, j := 0, len(height)-1
	for i < j {
		hi, hj := height[i], height[j]
		h := min(hi, hj)
		w := j - i
		res = max(res, h*w)
		if hi < hj {
			for i < j && height[i] <= hi {
				i++
			}
		} else {
			for i < j && height[j] <= hj {
				j--
			}
		}
	}
	return res
}
