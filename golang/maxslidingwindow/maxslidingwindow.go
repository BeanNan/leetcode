package maxslidingwindow

func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 || k < 1 {
		return nil
	}
	var window []int
	var result []int

	for index, value := range nums {
		if index >= k && window[0] <= index-k {
			window = window[1:]
		}

		for len(window) != 0 && nums[window[len(window)-1]] <= value {
			window = window[0 : len(window)-1]
		}

		window = append(window, index)

		if index >= k-1 {
			result = append(result, nums[window[0]])
		}
	}

	return result

}
